package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nihilKnight/grpc-section/plcide"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	client := pb.NewPLCRuntimeServiceClient(conn)

	// 订阅实时日志
	streamReq := &pb.LogSubscription{
		LevelFilter: pb.LogSubscription_ERROR,
	}
	stream, _ := client.StreamLogs(context.Background(), streamReq)
	go func() {
		for {
			logEntry, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[%s][%s] %s\n", 
				logEntry.Timestamp.AsTime().Format("2006-01-02 15:04:05"),
				logEntry.Level.String(), 
				logEntry.Message)
		}
	}()

	// 导出日志
	exportReq := &pb.LogExportRequest{
		Format:     "CSV",
		ChunkSize:  1024,
	}
	exportStream, _ := client.ExportLogs(context.Background(), exportReq)
	exportedData := []byte{}
	for {
		chunk, err := exportStream.Recv()
		if err != nil {
			break
		}
		exportedData = append(exportedData, chunk.Data...)
	}
	fmt.Println("导出日志内容:\n", string(exportedData))
}