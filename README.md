# go-protobug-example

### quick start

```bash
# Upstart server
$ dep ensure
$ realize start

# Setup client(https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)
$ git clone https://github.com/grpc/grpc
$ cd grpc
$ git submodule update --init
$ sudo apt-get install libgflags-dev
$ make grpc_cli
# for Ubuntu
```

### development setup

```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go get -u google.golang.org/grpc
$ sudo apt-get install libprotobuf-dev libprotoc-dev protobuf-compiler
```

### generates interfaces

```bash
$ protoc --go_out=plugins=grpc:./pb -I=./proto ./proto/user.proto
```