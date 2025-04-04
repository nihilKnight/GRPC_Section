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
    fmt.Printf("收到日志导出请求，chunk_size: %d\n", req.ChunkSize)
    // 模拟日志数据
    logEntry := &pb.LogEntry{
        Timestamp: &timestamp.Timestamp{Seconds: 1643723900, Nanos: 0},
        Level:     pb.LogEntry_INFO,
        Component: "模拟组件",
        Message:   "这是一个模拟日志消息",
    }
    return logEntry, nil
}

// CreateProject 实现项目创建功能
func (s *PLCRuntimeServer) CreateProject(ctx context.Context, req *pb.ProjectMetadata) (*pb.ProjectCreationResponse, error) {
    fmt.Printf("收到项目创建请求，项目信息:\n项目ID: %d\n名称: %s\n描述: %s\n版本字符串: %s\n任务数量: %d\n设备数量: %d\n",
        req.ProjectId, req.Name, req.Description, req.VersionString, req.TaskNum, req.DeviceNum)
    for _, task := range req.Tasks {
        fmt.Printf("任务ID: %d, 名称: %s, 类型: %d, 优先级: %d, 周期: %d, 触发器: %s\n",
            task.TaskId, task.Name, task.Type, task.Priority, task.Cycle, task.Trigger)
    }
    for _, device := range req.Devices {
        fmt.Printf("设备ID: %d, 类型: %d, 周期: %d\n", device.DeviceId, device.Type, device.Cycle)
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
    fmt.Printf("收到程序导入请求，程序ID: %d, 名称: %s, 内容:\n%s\n", req.ProgramId, req.ProgramName, req.Content)
    // 返回导入状态
    return &pb.ImportStatus{State: pb.ImportStatus_APPLIED}, nil
}

// ExportProgram 实现程序导出功能
func (s *PLCRuntimeServer) ExportProgram(ctx context.Context, req *pb.ProjectExportRequest) (*pb.ProjectMetadata, error) {
    fmt.Printf("收到程序导出请求，版本ID: %d, 版本字符串: %s, chunk_size: %d\n", req.VersionId, req.VersionString, req.ChunkSize)
    // 返回导出结果
    return &pb.ProjectMetadata{
        ProjectId:   1,
        Name:        "示例项目",
        Description: "这是一个示例项目",
        VersionString: "v1.0",
        TaskNum:     2,
        DeviceNum:   3,
        Tasks: []*pb.TaskMetadata{
            {TaskId: 1, Name: "任务1", Type: pb.TaskMetadata_RISING, Priority: 1, Cycle: 10, Trigger: "触发器1"},
            {TaskId: 2, Name: "任务2", Type: pb.TaskMetadata_FALLING, Priority: 2, Cycle: 20, Trigger: "触发器2"},
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
    log.Println("gRPC 服务器启动，监听端口：50051")
    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
