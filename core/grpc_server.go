package core

import (
	"bytes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "google.golang.org/grpc/encoding/gzip"
)

// 2M
const maxFileSize = 1 << 21

type Server interface {
	Listen() (err error)
	Close()
}

type ServerGRPC struct {
	fileStore   FileStore
	logger      zerolog.Logger
	server      *grpc.Server
	port        int
	certificate string
	key         string
}

type ServerGRPCConfig struct {
	Certificate string
	Key         string
	Port        int
}

func NewServerGRPC(cfg ServerGRPCConfig, fileStore FileStore) (s ServerGRPC, err error) {
	s.logger = zerolog.New(os.Stdout).With().Str("from", "server").Logger()

	if cfg.Port == 0 {
		err = errors.Errorf("Port must be specified")
		return
	}

	s.port = cfg.Port
	s.certificate = cfg.Certificate
	s.key = cfg.Key
	s.fileStore = fileStore
	return
}

func (s *ServerGRPC) Listen() (err error) {
	var (
		listener  net.Listener
		grpcOpts  []grpc.ServerOption
		grpcCreds credentials.TransportCredentials
	)

	listener, err = net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		err = errors.Wrapf(err, "failed to listen on port %d", s.port)
		return
	}

	if s.certificate != "" && s.key != "" {
		grpcCreds, err = credentials.NewServerTLSFromFile(s.certificate, s.key)
		if err != nil {
			err = errors.Wrapf(err, "failed to create tls grpc serve using cert %s and key %s", s.certificate, s.key)
			return
		}

		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	s.server = grpc.NewServer(grpcOpts...)
	RegisterGuploadServiceServer(s.server, s)

	err = s.server.Serve(listener)
	if err != nil {
		err = errors.Wrapf(err, "errored listening for grpc connections")
		return
	}
	return
}

func (s *ServerGRPC) Upload(stream GuploadService_UploadServer) (err error) {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}
	fileId := req.GetInfo().GetFileId()
	fileType := req.GetInfo().GetFileType()
	log.Printf("receive an upload request for fileId %s with type %s", fileId, fileType)

	data := bytes.Buffer{}
	filesize := 0

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				goto END
			}

			err = errors.Wrapf(err, "failed unexpectedly while reading chunks from stream")
			return err
		}
		chunk := req.GetContent()
		size := len(chunk)
		log.Printf("received a chunk with size: %d", size)

		filesize += size
		if filesize > maxFileSize {
			return logError(status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", filesize, maxFileSize))
		}

		_, err = data.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

END:
	_, err = s.fileStore.Save(fileId, fileType, data)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save file: %v", err))
	}

	err = stream.SendAndClose(&UploadStatus{
		Message: "Upload received with success",
		Code:    UploadStatusCode_Ok,
	})
	if err != nil {
		err = errors.Wrapf(err, "failed to send status code")
		return
	}
	s.logger.Info().Msg("file saved: " + fileId)
	return
}

func (s *ServerGRPC) Close() {
	if s.server != nil {
		s.server.Stop()
	}
	return
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
