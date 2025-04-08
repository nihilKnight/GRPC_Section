package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/nihilKnight/grpc-section/gen"
)

// PLCRuntimeServer 是 PLCRuntimeService 的实现
type PLCRuntimeServer struct{
    pb.UnimplementedPLCRuntimeServiceServer
}

// ExportLogs 实现日志导出功能
func (s *PLCRuntimeServer) ExportLogs(ctx context.Context, req *pb.LogExportRequest) (*pb.LogEntry, error) {
    fmt.Printf("Received Exporting Logs Request，chunk_size: %d\n", req.ChunkSize)
    // 模拟日志数据
    logEntry := &pb.LogEntry{
        Timestamp: &timestamp.Timestamp{Seconds: 1643723900, Nanos: 0},
        Level:     pb.LogEntry_INFO,
        Component: "Simulation Component",
        Message:   "This is a piece of simulation log.",
    }
    return logEntry, nil
}

// CreateProject 实现项目创建功能
func (s *PLCRuntimeServer) CreateProject(ctx context.Context, req *pb.ProjectMetadata) (*pb.ProjectCreationResponse, error) {
    fmt.Printf("Received Creating Project Request, Project Info:\nProject ID: %d\nName: %s\nDescription: %s\nVersion-String: %s\nTask Number: %d\nDevice Number: %d\n",
        req.ProjectId, req.Name, req.Description, req.VersionString, req.TaskNum, req.DeviceNum)
    for _, task := range req.Tasks {
        fmt.Printf("Task ID: %d, Name: %s, Type: %d, Priority: %d, Cycle: %d, Trigger: %s\n",
            task.TaskId, task.Name, task.Type, task.Priority, task.Cycle, task.Trigger)
    }
    for _, device := range req.Devices {
        fmt.Printf("Device ID: %d, Type: %d, Cycle: %d\n", device.DeviceId, device.Type, device.Cycle)
    }
    // 返回创建结果
    return &pb.ProjectCreationResponse{
        ProjectId:   req.ProjectId,
        CreatedAt:   &timestamp.Timestamp{Seconds: 1643723900, Nanos: 0},
        InitialVersion: "v1.0",
    }, nil
}

// ImportProgram 实现程序导入功能
func (s *PLCRuntimeServer) ImportProgram(ctx context.Context, req *pb.ProgramSource) (*pb.ImportStatus, error) {
    fmt.Printf("Received Importing Program Requeust，Program ID: %d, Name: %s, Content:\n%s\n", req.ProgramId, req.ProgramName, req.Content)
    // 返回导入状态
    return &pb.ImportStatus{State: pb.ImportStatus_APPLIED}, nil
}

// ExportProject 实现程序导出功能
func (s *PLCRuntimeServer) ExportProject(ctx context.Context, req *pb.ProjectExportRequest) (*pb.ProjectMetadata, error) {
    fmt.Printf("Received Exporting Project Requeust，Version ID: %d, Version String: %s, Chunk Size: %d\n", req.VersionId, req.VersionString, req.ChunkSize)
    // 返回导出结果
    return &pb.ProjectMetadata{
        ProjectId:   1,
        Name:        "Example Project",
        Description: "This is a simulation project.",
        VersionString: "v1.0",
        TaskNum:     2,
        DeviceNum:   3,
        Tasks: []*pb.TaskMetadata{
            {TaskId: 1, Name: "Task1", Type: pb.TaskMetadata_RISING, Priority: 1, Cycle: 10, Trigger: "Trigger1"},
            {TaskId: 2, Name: "Task2", Type: pb.TaskMetadata_FALLING, Priority: 2, Cycle: 20, Trigger: "Trigger2"},
        },
        Devices: []*pb.BusDeviceMetadata{
            {DeviceId: 1, Type: pb.BusDeviceMetadata_GPIO, Cycle: 10},
            {DeviceId: 2, Type: pb.BusDeviceMetadata_MODBUS, Cycle: 20},
            {DeviceId: 3, Type: pb.BusDeviceMetadata_OTHER, Cycle: 30},
        },
    }, nil
}

func main() {
    // 监听端口
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    // 创建 gRPC 服务器
    s := grpc.NewServer()
    // 注册服务
    pb.RegisterPLCRuntimeServiceServer(s, &PLCRuntimeServer{})
    log.Println("gRPC Server Started, listening on: 50051...")
    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
