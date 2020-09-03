module github.com/rtang03/grpc-server/cmd

go 1.15

replace github.com/rtang03/grpc-server/api => ./../api

replace github.com/rtang03/grpc-server/core => ./../core

require (
	github.com/rtang03/grpc-server/core v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/net v0.0.0-20200822124328-c89045814202
)
