# go-protobug-example

### quick start

```
$ dep ensure
$ realize start
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