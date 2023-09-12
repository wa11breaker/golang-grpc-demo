# Introduction

This project is an experimental attempt to reduce the server to server sending latency between two microservices under 3ms. 


## Setup

Before you can run the microservices, you'll need to set up some dependencies and generate the necessary code.

### Install `protoc-gen-go`

You need to install the `protoc-gen-go` plugin to generate Go code from protocol buffer definitions.

```bash
brew install protoc-gen-go
```

### Generate Go Code from Protobuf

Run the following commands to generate Go code from the protobuf definition files.

```bash
protoc -I=. --go_out=. --go-grpc_out=. cart-ms/proto/cart.proto
protoc -I=. --go_out=. --go-grpc_out=. discount-ms/proto/discount.proto
```

### Start Cart Microservice

Navigate to the `cart-ms` directory and run the Cart microservice server.

```bash
cd cart-ms
go run server.go
```

### Start Discount Microservice

Navigate to the `discount-ms` directory and run the Discount microservice server.

```bash
cd discount-ms
go run server.go
```
