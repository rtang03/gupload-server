package cmd

import (
	"github.com/rtang03/grpc-server/core"
	"github.com/urfave/cli/v2"
)

var Serve = cli.Command{
	Name:   "serve",
	Usage:  "initiates a gRPC server",
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
		server      core.Server
	)

	grpcServer, err := core.NewServerGRPC(core.ServerGRPCConfig{
		Port:        port,
		Certificate: certificate,
		Key:         key,
	})
	must(err)
	server = &grpcServer

	err = server.Listen()
	must(err)
	defer server.Close()

	return
}
