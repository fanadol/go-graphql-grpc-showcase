.PHONY: rungql rungrpc

rungql:
	go run ./cmd/graphql/main.go

rungrpc:
	go run ./cmd/grpc/main.go