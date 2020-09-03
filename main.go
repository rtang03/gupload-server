package main

import (
	"os"

  "github.com/rtang03/grpc-server/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "gupload",
		Usage:    "upload files",
		Commands: []*cli.Command{
		  &cmd.Serve,
		  &cmd.Upload,
    },
    Flags: []cli.Flag{
		  &cli.BoolFlag{
		    Name: "debug",
		    Usage: "enables debug logging",
      },
    },
	}

	app.Run(os.Args)
}
