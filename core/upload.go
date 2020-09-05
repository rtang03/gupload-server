package core

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
)

var UploadCommand = cli.Command{
	Name:   "upload",
	Usage:  "upload a file (max 4MB per file)",
	Action: uploadAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Value: "localhost:1313",
			Usage: "address of the server to connect to",
		},
		&cli.IntFlag{
			Name:  "chunk-size",
			Usage: "size of chunk messages (grpc only)",
			Value: 1 << 12,
		},
		&cli.StringFlag{
			Name:  "file",
			Usage: "file to upload",
		},
		&cli.StringFlag{
			Name:  "cacert",
			Usage: "path of a certifcate to add to the root CAs",
		},
		&cli.BoolFlag{
			Name:  "compress",
			Usage: "whether or not to enable payload compression",
		},
		&cli.StringFlag{
			Name:  "servername-override",
			Usage: "use serverNameOverride for tls ca cert",
		},
		&cli.StringFlag{
			Name:  "filename",
			Usage: "filename after upload",
		},
		&cli.StringFlag{
			Name:  "label",
			Usage: "label can be considered your organization id, e.g. org100",
		},
	},
}

func uploadAction(c *cli.Context) (err error) {
	var (
		chunkSize          = c.Int("chunk-size")
		address            = c.String("address")
		file               = c.String("file")
		rootCertificate    = c.String("cacert")
		compress           = c.Bool("compress")
		serverNameOverride = c.String("servername-override")
		filename           = c.String("filename")
		mspid              = c.String("label")
		client             Client
	)

	if address == "" {
		must(errors.New("address"))
	}

	if file == "" {
		must(errors.New("file must be set"))
	}

	grpcClient, err := NewClientGRPC(ClientGRPCConfig{
		Address:            address,
		RootCertificate:    rootCertificate,
		Compress:           compress,
		ChunkSize:          chunkSize,
		ServerNameOverride: serverNameOverride,
		Filename:           filename,
		Mspid:              mspid,
	})
	must(err)
	client = &grpcClient

	stat, err := client.UploadFile(context.Background(), file)
	must(err)
	defer client.Close()

	fmt.Printf("⏱  Time duration (ms): %d\n", stat.FinishedAt.Sub(stat.StartedAt).Milliseconds())

	return
}
