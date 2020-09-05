package core

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var ServeCommand = cli.Command{
	Name:   "serve",
	Usage:  "initiates a gRPC upload server (max 4MB per file)",
	Action: serveAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "port",
			Usage: "port to bind to",
			Value: 1313,
		},
		&cli.StringFlag{
			Name:  "key",
			Usage: "path to TLS key",
		},
		&cli.StringFlag{
			Name:  "certificate",
			Usage: "path to TLS certificate",
		},
	},
}

func serveAction(c *cli.Context) (err error) {
	var (
		port        = c.Int("port")
		key         = c.String("key")
		certificate = c.String("certificate")
		server      Server
	)
	fileStore := NewDiskStore("uploaded")

	grpcServer, err := NewServerGRPC(ServerGRPCConfig{
		Port:        port,
		Certificate: certificate,
		Key:         key,
	}, fileStore)
	must(err)
	server = &grpcServer

	fmt.Printf("🚀 Gupload server listen at: %d\n", port)
	err = server.Listen()
	must(err)
	defer server.Close()
	return
}
