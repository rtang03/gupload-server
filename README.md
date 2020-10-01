## Overview
This small utility setup server/cli: "upload-only" ftp-like server; with TLS + grpc transport.

### Motivation
When doing multi-cloud deployment of Hyperledger Fabric, peers of different organizations run on different cloud providers.
It needs a mechanism to share tls root certs, and/or crypto material, as an out-of-band communication process. This utility
is created as Pod, besides "peer" pod. It enables the out-of-band file exchange process uses the same networking transport
of inter-peer communications. A max filesize is 4 MB. TLS is required for SNI-based routing.

### Pre-requisite
- Go v1.15 +
- [Protocol buffer compiler](https://grpc.io/docs/languages/go/quickstart/)
- [Golang editor](https://jaxenter.com/top-5-ides-go-146348.html)

### Instructions
```text
NAME:
   gupload - Upload and download files with grpcs

USAGE:
   gupload [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   serve     initiates a gRPC upload server (max 4MB per file)
   upload    upload a file (max 4MB per file)
   download  download a file
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

Use `serve` to initiate a `gRPC` server and `upload` to upload a file to a given address.

### Serve GRPC

`grpc` is the default mechanism used for both clients and servers.

There are two forms of running it:

- via TLS-based http2

To use tls-based connections:

```shell script
# Create a server
./build/gupload serve --key ./cert/tls.key --certificate ./cert/tls.crt

# When doing local development with above cert/key pair;
# see this issue https://github.com/golang/go/issues/39568
# if we use localhost in the tls cert for local dev, need to set below env
# this workaround may later break, for golang version beyong v1.15
export GODEBUG=x509ignoreCN=0
```

### Upload a file
```shell script
# Upload a file: with mandatory fields
./build/gupload upload \
    --cacert ./cert/tls.crt \
    --infile main.go \
    --public \
    --outfile main.go \
    --address localhost:1313

```
If `public` flag is false, the uploaded filename will be placed at `fileserver` directory; its filename will be `main.go`. Or
otherwise, the uploaded file will be sent to `fileserver/public` directory in the server.

Note that `gupload` client is not able to download non-public files.

The default address is `localhost:1313`.

Also, can use `--servername-override`, when TLS is enabled.

### Download a file
```shell script
./build/gupload download \
    --cacert ./cert/tls.crt \
    --file test.txt \
    --address localhost:1313
```

It will download file from `fileserver/public` directory.


### Credits
The tool is adapted from:
- https://github.com/cirocosta/gupload
- https://github.com/techschool/pcbook-go

### Reference Info
- [protobuff for go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [go-grpc-tutorial](https://tutorialedge.net/golang/go-grpc-beginners-tutorial/)
- [youtube tutorial #1](https://www.youtube.com/watch?v=BdzYdN_Zd9Q)
- [youtube tutorial #2](https://www.youtube.com/watch?v=i2p0Snwk4gc)
- [example 1](https://gitlab.com/pantomath-io/demo-grpc)
- [example 2](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)
- [publish to gh registry](https://github.com/actions/starter-workflows/blob/main/ci/docker-publish.yml)
- [example 3](https://github.com/fuskovic/server-streaming)

### Development
```shell script
# generate protocol buffers
protoc --proto_path=core --go_out=plugins=grpc:core --go_opt=paths=source_relative core/service.proto

# compile
go build -i -v -o build/gupload main.go

# to trigger the docker image creation and send to Github Container Registry
git tag v0.0.2

git push origin v0.0.2
```
