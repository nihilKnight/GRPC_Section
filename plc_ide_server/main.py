import grpc
from concurrent import futures
import time
import hashlib
from google.protobuf import timestamp_pb2
from google.protobuf.empty_pb2 import Empty
from plc_ide import plc_ide_pb2 as pb
from plc_ide import plc_ide_pb2_grpc as pb_grpc

class PLCRuntimeService(pb_grpc.PLCRuntimeServiceServicer):
    def __init__(self):
        # 初始化模拟数据存储
        self.projects = {}
        self.config_versions = {}
        self.log_buffer = []

    # ------------------ 日志管理服务 ------------------
    def StreamLogs(self, request, context):
        """实时日志流订阅"""
        print(f"[服务端] 收到日志订阅请求: 级别={pb.LogSubscription.LogLevel.Name(request.level_filter)}")
        
        # 模拟持续生成日志
        while True:
            log_entry = pb.LogEntry(
                timestamp=timestamp_pb2.Timestamp(seconds=int(time.time())),
                level=request.level_filter,
                source="PLC.Runtime",
                message="模拟日志消息",
                context={"device_id": "PLC-001"}
            )
            yield log_entry
            time.sleep(1)

    def ExportLogs(self, request, context):
        """日志导出分块传输"""
        print(f"[服务端] 收到日志导出请求: 格式={request.format}")
        
        # 生成模拟CSV数据
        csv_data = "timestamp,level,message\n"
        for _ in range(3):  # 生成3行示例数据
            csv_data += f"{int(time.time())},ERROR,这是一个错误日志\n"
        
        # 分块发送
        chunk = pb.LogChunk(
            data=csv_data.encode(),
            sequence=1,
            checksum=hashlib.sha256(csv_data.encode()).hexdigest()
        )
        yield chunk

    # ------------------ 配置管理服务 ------------------
    def ImportProjectConfig(self, request_iterator, context):
        """配置导入（客户端流式）"""
        print("[服务端] 开始接收配置导入...")
        config_data = bytearray()
        total_chunks = 0
        
        try:
            for chunk in request_iterator:
                print(f"接收分块 {chunk.current_chunk}/{chunk.total_chunks}")
                config_data.extend(chunk.content)
                total_chunks = chunk.total_chunks
            
            # 验证配置完整性
            if len(config_data) == 0:
                return pb.ImportStatus(state=pb.ImportStatus.State.FAILED, error_detail="空配置")
            
            # 存储配置版本
            version_id = f"v{len(self.config_versions)+1}"
            self.config_versions[version_id] = config_data
            
            return pb.ImportStatus(
                state=pb.ImportStatus.State.APPLIED,
                progress=1.0,
                initial_version=version_id
            )
        except Exception as e:
            return pb.ImportStatus(state=pb.ImportStatus.State.FAILED, error_detail=str(e))

    def ExportProjectConfig(self, request, context):
        """配置导出（服务端流式）"""
        print(f"[服务端] 导出项目配置: {request.version_id}")
        
        if request.version_id not in self.config_versions:
            context.abort(grpc.StatusCode.NOT_FOUND, "版本不存在")
        
        data = self.config_versions[request.version_id]
        chunk_size = 1024  # 1KB分块
        
        for i in range(0, len(data), chunk_size):
            yield pb.ConfigChunk(
                content=data[i:i+chunk_size],
                total_chunks=len(data)//chunk_size + 1,
                current_chunk=i//chunk_size + 1,
                config_hash=hashlib.sha256(data).hexdigest()
            )

    # ------------------ 项目管理服务 ------------------
    def CreateProject(self, request, context):
        """创建新项目"""
        print(f"[服务端] 创建项目: {request.name}")
        
        if request.project_id in self.projects:
            context.abort(grpc.StatusCode.ALREADY_EXISTS, "项目已存在")
        
        new_project = {
            "metadata": request,
            "versions": []
        }
        self.projects[request.project_id] = new_project
        
        return pb.ProjectCreationResponse(
            project_id=request.project_id,
            created_at=timestamp_pb2.Timestamp(seconds=int(time.time())),
            initial_version="v1.0"
        )

    def ListProjects(self, request, context):
        """列出所有项目"""
        print(f"[服务端] 查询项目列表，过滤条件: {request.name_pattern}")
        
        project_list = pb.ProjectList()
        for pid, project in self.projects.items():
            if request.name_pattern and request.name_pattern not in project["metadata"].name:
                continue
            
            summary = pb.ProjectList.ProjectSummary(
                project_id=pid,
                name=project["metadata"].name,
                last_modified=timestamp_pb2.Timestamp(seconds=int(time.time())),
                version_count=len(project["versions"])
            )
            project_list.projects.append(summary)
        
        return project_list

    # ------------------ 其他服务方法 ------------------
    def ListProjectVersions(self, request, context):
        """列出项目版本"""
        if request.project_id not in self.projects:
            context.abort(grpc.StatusCode.NOT_FOUND, "项目不存在")
        
        versions = pb.ProjectVersions()
        versions.versions.extend([
            pb.ProjectVersions.VersionInfo(
                id=f"v{i+1}",\
                created_at=timestamp_pb2.Timestamp(seconds=int(time.time())),\
                author="system",\
                comment="自动生成版本"
            ) for i in range(3)  # 模拟3个版本
        ])
        return versions

    def DeleteProject(self, request, context):
        """删除项目"""
        if request.project_id not in self.projects:
            return pb.OperationStatus(success=False, error_code="NOT_FOUND")
        
        del self.projects[request.project_id]
        return pb.OperationStatus(success=True)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb_grpc.add_PLCRuntimeServiceServicer_to_server(PLCRuntimeService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("服务端已启动，监听端口 50051...")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()