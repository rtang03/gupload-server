https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
https://www.youtube.com/watch?v=BdzYdN_Zd9Q
https://www.youtube.com/watch?v=i2p0Snwk4gc
https://gitlab.com/pantomath-io/demo-grpc
https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2
https://github.com/techschool/pcbook-go/blob/master/cmd/server/main.go

protoc --proto_path=src --go_out=src --go_opt=paths=source_relative  src/api.proto
protoc --proto_path=api --go_out=plugins=grpc:api --go_opt=paths=source_relative api/service.proto
protoc --proto_path=core --go_out=plugins=grpc:core --go_opt=paths=source_relative core/service.proto


export GODEBUG=x509ignoreCN=0
https://github.com/golang/go/issues/39568
```shell script
go build -v -i -o gupload main.go

./gupload serve --key ./cert/server.key --certificate ./cert/server.crt

./gupload upload --file [any-file] --chunk-size 4096 --cacert ./grpc-server/cert/server.crt --address 127.0.0.1:1313
```
