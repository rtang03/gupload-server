package core

import (
	"bytes"
	"fmt"
	"google.golang.org/grpc/codes"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "google.golang.org/grpc/encoding/gzip"
)

type Client interface {
	UploadFile(ctx context.Context, f string) (stats Stats, err error)
	DownloadFile(f string) (err error)
	Check(ctx context.Context, label string, counter int) (pingStats PingStats, err error)
	Close()
}

type ClientGRPC struct {
	conn            *grpc.ClientConn
	client          GuploadServiceClient
	chunkSize       int
	filename        string
	usePublicFolder bool
}

type ClientGRPCConfig struct {
	Address            string
	RootCertificate    string
	Compress           bool
	ServerNameOverride string
	Filename           string
	UsePublicFolder    bool
}

func NewClientGRPC(cfg ClientGRPCConfig) (c ClientGRPC, err error) {
	var (
		grpcOpts  []grpc.DialOption
		grpcCreds credentials.TransportCredentials
	)
	// 4096 fixed
	c.chunkSize = 1 << 12
	c.usePublicFolder = cfg.UsePublicFolder
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
		// for use in health_check
		grpcOpts = append(grpcOpts, grpc.WithInsecure())

		// comment this, because non-ssl operation is guarded by CLI check. This check is reducdancy. Can remove it later when things go well
		//err = errors.Errorf("non-ssl operation is not suported")
		//return
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
		writing  = true
		buf      []byte
		n        int
		file     *os.File
		status   *UploadStatus
		fileType string
	)

	fi, err := os.Stat(f)
	if err != nil {
		err = errors.Wrapf(err, "failed to find file %s", f)
		return
	}

	if fi.Size() > maxFileSize {
		fmt.Println("Too big file size to send (max 4MB)")
		err = errors.Wrapf(err, "too big file size to send (max 4M)")
		return
	}

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
	if c.usePublicFolder == true {
		fileType = "public"
	} else {
		fileType = "private"
	}
	req := &Chunk{
		Data: &Chunk_Info{
			Info: &UploadFileInfo{
				Filename: c.filename,
				FileType: fileType,
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

	if status.Code != StatusCode_Ok {
		err = errors.Errorf("upload filed - msg: %s", status.Message)
		return
	}
	return
}

func (c *ClientGRPC) DownloadFile(fileName string) (err error) {
	req := &FileRequest{
		Filename: fileName,
	}
	stream, err := c.client.Download(context.Background(), req)
	if err != nil {
		return err
	}

	var downloaded int64
	var buffer bytes.Buffer

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			if err := ioutil.WriteFile(fileName, buffer.Bytes(), 0644); err != nil {
				return err
			}
			break
		}
		if err != nil {
			buffer.Reset()
			return err
		}
		shard := res.GetShard()
		shardSize := len(shard)
		downloaded += int64(shardSize)

		buffer.Write(shard)
		fmt.Printf("\r%s", strings.Repeat(" ", 25))
		fmt.Printf("\r%s downloaded", humanize.Bytes(uint64(downloaded)))
	}

	return nil
}

func (c *ClientGRPC) Check(ctx context.Context, label string, counter int) (pingStats PingStats, err error) {
	var res *HealthCheckResponse
	req := new(HealthCheckRequest)

	pingStats.pingStartAt = time.Now()
	req.PingAt = time.Now().UTC().String()
	req.Label = label
	req.Counter = strconv.Itoa(counter)

	res, err = c.client.Check(ctx, req)

	pingStats.pingFinishedAt = time.Now()

	pingStats.ok = false

	if err == nil {
		pingStats.serverReceivedAt = res.GetReceivedAt()

		if res.GetStatus() == HealthCheckResponse_SERVING {
			pingStats.ok = true
			return pingStats, nil
		}
		return pingStats, nil
	}

	switch grpc.Code(err) {
	case
		codes.Aborted,
		codes.DataLoss,
		codes.DeadlineExceeded,
		codes.Internal,
		codes.Unavailable:
		// non-fatal errors
	default:
		return pingStats, err
	}

	return pingStats, err
}

func (c *ClientGRPC) Close() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
