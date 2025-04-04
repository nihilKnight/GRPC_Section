package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetConnection() *grpc.ClientConn {
    conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}