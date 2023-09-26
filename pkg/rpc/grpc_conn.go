// Package rpc provides a gRPC client.
package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcConn creates a gRPC client.
func GrpcConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("query:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
