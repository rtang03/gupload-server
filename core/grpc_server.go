package core

import (
	"io"
	"net"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rtang03/grpc-server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "google.golang.org/grpc/encoding/gzip"
)

type Server interface {
	Listen() (err error)
	Close()
}

type ServerGRPC struct {
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

func NewServerGRPC(cfg ServerGRPCConfig) (s ServerGRPC, err error) {
	s.logger = zerolog.New(os.Stdout).With().Str("from", "server").Logger()

	if cfg.Port == 0 {
		err = errors.Errorf("Port must be specified")
		return
	}

	s.port = cfg.Port
	s.certificate = cfg.Certificate
	s.key = cfg.Key

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
	api.RegisterGuploadServiceServer(s.server, s)

	err = s.server.Serve(listener)
	if err != nil {
		err = errors.Wrapf(err, "errored listening for grpc connections")
		return
	}
	return
}

func (s *ServerGRPC) Upload(stream api.GuploadService_UploadServer) (err error) {
	for {
		_, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				goto END
			}

			err = errors.Wrapf(err, "failed unexpectedly while reading chunks from stream")
			return
		}
	}

	// s.logger.Info().Msg("upload received")
END:
	err = stream.SendAndClose(&api.UploadStatus{
		Message: "Upload received with success",
		Code:    api.UploadStatusCode_Ok,
	})
	if err != nil {
		err = errors.Wrapf(err, "failed to send status code")
		return
	}
	return
}

func (s *ServerGRPC) Close() {
	if s.server != nil {
		s.server.Stop()
	}
	return
}
