package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"

	"github.com/rtang03/grpc-server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

// authenticateAgent check the client credentials
func authenticateClient(ctx context.Context, s *api.Server) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		clientLogin := strings.Join(md["login"], "")
		clientPassword := strings.Join(md["password"], "")

		if clientLogin != "john" {
			return "", fmt.Errorf("unknown user %s", clientLogin)
		}

		if clientPassword != "doe" {
			return "", fmt.Errorf("bad password %s", clientPassword)
		}

		log.Printf("authenticated client: %s", clientLogin)
		return "42", nil
	}

	return "", fmt.Errorf("missing credentials")
}

// unaryInterceptor calls authenticateClient with current context
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s, ok := info.Server.(*api.Server)
	if !ok {
		return nil, fmt.Errorf("unable to cast server")
	}
	clientID, err := authenticateClient(ctx, s)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, clientIDKey, clientID)
	return handler(ctx, req)
}

// main start a gRPC server and waits for connection
func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	// create a listener on TCP port 9000
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := api.Server{}

	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor)}
	grpcServer := grpc.NewServer(opts...)

	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}

}
