package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "gupload",
		Usage:    "upload files",
		Commands: []*cli.Command{},
	}

	app.Run(os.Args)
}
