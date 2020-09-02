package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rtang03/grpc-server/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Authetnication struct {
	Login    string
	Password string
}

// GetRequestMetadata gets the current request metadata
func (a *Authetnication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authetnication) RequireTransportSecurity() bool {
	return true
}

func main() {
	fmt.Println("Env Variables")
	var password string
	password = os.Getenv("PASS")
	fmt.Printf("DB Password: %s\n", password)

	var conn *grpc.ClientConn

	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	auth := Authetnication{
		Login:    "john",
		Password: "doe",
	}

	conn, err = grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
