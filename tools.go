//go:build never
// +build never

package tools

import (
	_ google.golang.org/protobuf
	_ google.golang.org/protobuf/cmd/protoc-gen-go
	_ google.golang.org/grpc/cmd/protoc-gen-go-grpc
	_ google.golang.org/grpc
	_ github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	_ github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
)
