# Graphql Go Example

This project contains single graphql server combined with gRPC to connect to internal service.

# Proto Compile

Used for generate golang file from `.proto` file.

```
protoc -I=. internal/user/proto/user.proto --go_out=plugins=grpc:.
```

# Prerequisites

- Golang 1.13
- protoc (Optional)

# How To Run

Run the graphql server

```
make rungql
```

Run the gRPC server

```
make rungrpc
```

Graphql playground will running on `port :8080`
