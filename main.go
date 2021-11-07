// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gopackages command is a diagnostic tool that demonstrates
// how to use golang.org/x/tools/go/packages to load, parse,
// type-check, and print one or more Go packages.
// Its precise output is unspecified and may change.
package main

import (
	"github.com/rtang03/grpc-server/core"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "gupload",
		Version: "0.1.10",
		Usage:   "Upload and download files with grpcs / grpc health check",
		Commands: []*cli.Command{
			&core.ServeCommand,
			&core.UploadCommand,
			&core.DownloadCommand,
			&core.HealthCheckCommand,
		},
	}

	_ = app.Run(os.Args)
}
