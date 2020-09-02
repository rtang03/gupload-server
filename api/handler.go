package api

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

// SayHello
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
