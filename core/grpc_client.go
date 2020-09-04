package core

import (
	"io"
	"os"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "google.golang.org/grpc/encoding/gzip"
)

type Client interface {
	UploadFile(ctx context.Context, f string) (stats Stats, err error)
	Close()
}

type ClientGRPC struct {
	conn      *grpc.ClientConn
	client    GuploadServiceClient
	chunkSize int
	filename  string
	mspid     string
}

type ClientGRPCConfig struct {
	Address            string
	ChunkSize          int
	RootCertificate    string
	Compress           bool
	ServerNameOverride string
	Filename           string
	Mspid              string
}

func NewClientGRPC(cfg ClientGRPCConfig) (c ClientGRPC, err error) {
	var (
		grpcOpts  []grpc.DialOption
		grpcCreds credentials.TransportCredentials
	)
	c.mspid = cfg.Mspid
	c.filename = cfg.Filename

	if cfg.Address == "" {
		err = errors.Errorf("address must be specified")
		return
	}

	if cfg.Compress {
		grpcOpts = append(grpcOpts, grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip")))
	}

	if cfg.RootCertificate != "" {
		grpcCreds, err = credentials.NewClientTLSFromFile(cfg.RootCertificate, cfg.ServerNameOverride)
		if err != nil {
			err = errors.Wrapf(err, "failed create grpc tls client via root-cert %s", cfg.RootCertificate)
			return
		}

		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(grpcCreds))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}

	switch {
	case cfg.ChunkSize == 0:
		err = errors.Errorf("ChunkSize must be specified")
		return
	case cfg.ChunkSize > (1 << 22):
		err = errors.Errorf("ChunkSize must be < than 4MB")
		return
	default:
		c.chunkSize = cfg.ChunkSize
	}

	c.conn, err = grpc.Dial(cfg.Address, grpcOpts...)
	if err != nil {
		err = errors.Wrapf(err, "failed to start grpc connection with address %s", cfg.Address)
		return
	}

	c.client = NewGuploadServiceClient(c.conn)

	return
}

func (c *ClientGRPC) UploadFile(ctx context.Context, f string) (stats Stats, err error) {
	var (
		writing = true
		buf     []byte
		n       int
		file    *os.File
		status  *UploadStatus
	)

	file, err = os.Open(f)
	if err != nil {
		err = errors.Wrapf(err, "failed to open file %s", f)
		return
	}
	defer file.Close()

	stream, err := c.client.Upload(ctx)
	if err != nil {
		err = errors.Wrapf(err, "failed to create upload stream for file %s", f)
		return
	}
	defer stream.CloseSend()

	stats.StartedAt = time.Now()

	// file info
	req := &Chunk{
		Data: &Chunk_Info{
			Info: &UploadFileInfo{
				FileId:   c.filename,
				FileType: c.mspid,
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		err = errors.Wrapf(err, "error send file header")
		return
	}

	// binary data
	buf = make([]byte, c.chunkSize)
	for writing {
		n, err = file.Read(buf)
		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}
			err = errors.Wrapf(err, "errored while copying from file to buf")
			return
		}

		err = stream.Send(&Chunk{
			Data: &Chunk_Content{
				Content: buf[:n],
			},
		})
		if err != nil {
			err = errors.Wrapf(err, "failed to send chunk via stream")
			return
		}
	}

	stats.FinishedAt = time.Now()

	status, err = stream.CloseAndRecv()
	if err != nil {
		err = errors.Wrapf(err, "failed to receive upstream status response")
		return
	}

	if status.Code != UploadStatusCode_Ok {
		err = errors.Errorf("upload filed - msg: %s", status.Message)
		return
	}
	return
}

func (c *ClientGRPC) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
