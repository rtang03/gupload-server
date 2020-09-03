https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
https://www.youtube.com/watch?v=BdzYdN_Zd9Q
https://www.youtube.com/watch?v=i2p0Snwk4gc
https://gitlab.com/pantomath-io/demo-grpc
https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2

protoc --proto_path=src --go_out=src --go_opt=paths=source_relative  src/api.proto
protoc --proto_path=api --go_out=plugins=grpc:api --go_opt=paths=source_relative api/service.proto

```shell script
go mod init github.com/rtang03/grpc-server/client
```
