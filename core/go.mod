module github.com/rtang03/grpc-server/core

go 1.15

replace github.com/rtang03/grpc-server/api => ./../api

require (
	github.com/pkg/errors v0.8.1
	github.com/rs/zerolog v1.19.0
	github.com/rtang03/grpc-server/api v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.31.1
)
