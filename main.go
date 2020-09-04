package main

import (
	"github.com/rtang03/grpc-server/core"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gupload",
		Usage: "gprc upload files",
		Commands: []*cli.Command{
			&core.ServeCommand,
			&core.UploadCommand,
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "enables debug logging",
			},
		},
	}

	app.Run(os.Args)
}
