// Package rpc provides a gRPC client.
package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GrpcConn creates a gRPC client.
func GrpcConn() (*grpc.ClientConn, error) {
	// Configure gRPC client to connect securely.
	creds := credentials.NewClientTLSFromCert(nil, "")

	conn, err := grpc.Dial("query:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
