package core

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "google.golang.org/grpc/encoding/gzip"
)

// 4M
const maxFileSize = 1 << 22

// 4096
// upload location: fileserver
// download location: fileserver/public
const filesDir = "fileserver/public"

type Server interface {
	Listen() (err error)
	Close()
}

type ServerGRPC struct {
	fileStore   FileStore
	server      *grpc.Server
	port        int
	certificate string
	key         string
	mu          sync.Mutex
	// statusMap stores the serving status of the services this Server monitors.
	statusMap map[string]HealthCheckResponse_ServingStatus
}

type ServerGRPCConfig struct {
	Certificate string
	Key         string
	Port        int
}

func NewServerGRPC(cfg ServerGRPCConfig, fileStore FileStore) (s ServerGRPC, err error) {

	if cfg.Port == 0 {
		err = errors.Errorf("Port must be specified")
		return
	}

	s.port = cfg.Port
	s.certificate = cfg.Certificate
	s.key = cfg.Key
	s.fileStore = fileStore

	// healthcheck
	s.statusMap = make(map[string]HealthCheckResponse_ServingStatus)

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

func (s *ServerGRPC) Download(request *FileRequest, stream GuploadService_DownloadServer) error {
	var shard []byte
	fileName := request.GetFilename()
	path := filepath.Join(filesDir, fileName)

	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return logError(status.Errorf(codes.NotFound, "filepath is invalid"))
	}
	fileSize := fileInfo.Size()

	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return logError(status.Errorf(codes.Aborted, "cannot open file"))
	}
	defer f.Close()

	var totalBytesStreamed int64

	for totalBytesStreamed < fileSize {
		bytesleft := fileSize - totalBytesStreamed
		if bytesleft < 1024 {
			shard = make([]byte, bytesleft)
		} else {
			shard = make([]byte, 1024)
		}
		bytesRead, err := f.Read(shard)
		if err == io.EOF {
			log.Println("download complete")
			break
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&FileResponse{
			Shard: shard,
		}); err != nil {
			return err
		}
		totalBytesStreamed += int64(bytesRead)
	}
	log.Println("download complete: " + fileName)
	return nil
}

func (s *ServerGRPC) Upload(stream GuploadService_UploadServer) (err error) {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive file info"))
	}
	fileId := req.GetInfo().GetFilename()
	fileType := req.GetInfo().GetFileType()
	log.Printf("receive an upload request for fileId '%s' with type '%s'", fileId, fileType)

	data := bytes.Buffer{}
	filesize := 0

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				log.Println("upload complete")
				break
			}

			err = errors.Wrapf(err, "failed unexpectedly while reading chunks from stream")
			return err
		}
		chunk := req.GetContent()
		size := len(chunk)
		fmt.Print("‣")

		filesize += size
		if filesize > maxFileSize {
			return logError(status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", filesize, maxFileSize))
		}

		_, err = data.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	_, err = s.fileStore.Save(fileId, fileType, data)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save file: %v", err))
	}

	err = stream.SendAndClose(&UploadStatus{
		Message: "Upload received with success",
		Code:    StatusCode_Ok,
	})
	if err != nil {
		err = errors.Wrapf(err, "failed to send status code")
		return
	}
	log.Printf("file saved (%s) : %s", fileType, fileId)
	return
}

func (s *ServerGRPC) Check(ctx context.Context, in *HealthCheckRequest) (*HealthCheckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("label:%s-%s ping at %s\n", in.Label, in.Counter, in.PingAt)

	if in.Service == "" {
		// check the server overall health status.
		return &HealthCheckResponse{
			Status:     HealthCheckResponse_SERVING,
			ReceivedAt: time.Now().UTC().String(),
		}, nil
	}

	if healthStatus, ok := s.statusMap[in.Service]; ok {
		return &HealthCheckResponse{
			Status:     healthStatus,
			ReceivedAt: time.Now().UTC().String(),
		}, nil
	}
	return nil, status.Error(codes.NotFound, "unknown service")
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
