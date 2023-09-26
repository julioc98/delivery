// Package rpc provides a gRPC client.
package rpc

import "google.golang.org/grpc"

// GrpcConn creates a gRPC client.
func GrpcConn() (*grpc.ClientConn, error) {
	// Connect to a server
	conn, err := grpc.Dial("query:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
