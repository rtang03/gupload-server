module github.com/rtang03/grpc-server

go 1.15

replace github.com/rtang03/grpc-server/core => ./core

require (
	github.com/rtang03/grpc-server/core v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.2.0
	google.golang.org/protobuf v1.25.0 // indirect
)
