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
		&cli.StringFlag{
			Name:  "infile",
			Usage: "local filename to upload",
		},
		&cli.StringFlag{
			Name:  "cacert",
			Usage: "path of a certifcate to add to the root CAs",
		},
		&cli.StringFlag{
			Name:  "servername-override",
			Usage: "use serverNameOverride for tls ca cert",
		},
		&cli.StringFlag{
			Name:  "outfile",
			Usage: "output filename after upload",
		},
		&cli.BoolFlag{
			Name:  "public",
			Usage: "send to public download folder",
			Value: false,
		},
	},
}

func uploadAction(c *cli.Context) (err error) {
	var (
		address            = c.String("address")
		infile             = c.String("infile")
		rootCertificate    = c.String("cacert")
		serverNameOverride = c.String("servername-override")
		outfile            = c.String("outfile")
		public             = c.Bool("public")
		client             Client
	)

	if address == "" {
		must(errors.New("address"))
	}

	if infile == "" {
		must(errors.New("infile must be set"))
	}

	if rootCertificate == "" {
		must(errors.New("cacert must be set"))
	}

	grpcClient, err := NewClientGRPC(ClientGRPCConfig{
		Address:            address,
		RootCertificate:    rootCertificate,
		Compress:           true,
		ServerNameOverride: serverNameOverride,
		Filename:           outfile,
		UsePublicFolder:    public,
	})
	must(err)
	client = &grpcClient

	stat, err := client.UploadFile(context.Background(), infile)
	must(err)
	defer client.Close()

	fmt.Printf("⏱  Time duration (ms): %d\n", stat.FinishedAt.Sub(stat.StartedAt).Milliseconds())

	return
}
