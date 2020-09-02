module server

go 1.15

replace github.com/rtang03/grpc-server/api => ../api

require (
	github.com/rtang03/grpc-server/api v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.31.1
)
