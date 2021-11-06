package core

import (
	"errors"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

var HealthCheckCommand = cli.Command{
	Name:   "ping",
	Usage:  "health check",
	Action: healtCheckAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Value: "localhost:1313",
			Usage: "address of the server to connect to",
		},
		&cli.StringFlag{
			Name:  "label",
			Value: "local",
			Usage: "Default: local. receiving server used to identify ping-client",
		},
		&cli.StringFlag{
			Name:  "interval",
			Value: "1s",
			Usage: "Default 1s. e.g. 1s, 5s, 100ms, 1h",
		},
		&cli.StringFlag{
			Name:  "cacert",
			Usage: "Deprecated. path of a certifcate to add to the root CAs",
		},
		&cli.StringFlag{
			Name:  "servername-override",
			Usage: "Deprecated. use serverNameOverride for tls ca cert",
		},
	},
}

func healtCheckAction(c *cli.Context) (err error) {
	var (
		address            = c.String("address")
		label              = c.String("label")
		serverNameOverride = c.String("servername-override")
		interval           = c.String("interval")
		client             Client
	)

	if address == "" {
		must(errors.New("address"))
	}

	duration, err := time.ParseDuration(interval)
	if err != nil {
		must(errors.New("invalid interval"))
	}

	grpcClient, err := NewClientGRPC(ClientGRPCConfig{
		Address:            address,
		RootCertificate:    "",
		Compress:           true,
		ServerNameOverride: serverNameOverride,
	})
	must(err)
	client = &grpcClient

	counter := 0

	for {
		counter++

		stats, err := client.Check(context.Background(), label, counter)

		if !stats.ok || err != nil {
			log.Printf("can't connect grpc server: %v, code: %v\n", err, grpc.Code(err))
		} else {
			log.Printf("label:%s-%d duration (ms) %d; --> %s\n", label, counter, stats.pingFinishedAt.Sub(stats.pingStartAt).Milliseconds(), stats.serverReceivedAt)
		}
		<-time.After(duration)
	}
}
