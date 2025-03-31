// plc_runtime_grpc_server.cc
#include "plc_ide_nrt_v1.grpc.pb.h"
#include "plc_ide_nrt_v1.pb.h"
#include "plc_runtime_grpc_server.h"
#include <grpcpp/grpcpp.h>
#include <iostream>
#include <queue>
#include <mutex>
#include <condition_variable>
#include <thread>
#include <chrono>

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using grpc::ServerWriter;
using plcide::LogEntry;
using plcide::LogExportRequest;

// 定义日志等级
enum class LogLevel {
    DEBUG = 0,
    INFO = 1,
    WARNING = 2,
    ERROR = 3,
    UNKNOWN = 4
};

class PLCRuntimeServiceImpl final : public plcide::PLCRuntimeService::Service {
public:
    PLCRuntimeServiceImpl() {}

    Status ExportLogs(ServerContext* context, const LogExportRequest* request, ServerWriter<LogEntry>* writer) override {
        // 处理请求并返回响应
        // 从日志队列中读取日志并发送
        std::unique_lock<std::mutex> lock(mutex_);
        while (true) {
            while (!log_queue_.empty()) {
                LogEntry log = log_queue_.front();
                log_queue_.pop();
                lock.unlock();
                writer->Write(log);
                lock.lock();
            }
            if (is_server_stopping_) {
                break;
            }
            cond_var_.wait(lock);
        }

        return Status::OK;
    }

    void AddLog(const std::string& component, const std::string& message, LogLevel level) {
        LogEntry log;
        log.set_component(component);
        log.set_message(message);
        log.set_level(static_cast<plcide::LogEntry_LogLevel>(level));
        auto timestamp = new google::protobuf::Timestamp();
        auto now = std::chrono::system_clock::now();
        auto duration = now.time_since_epoch();
        auto seconds = std::chrono::duration_cast<std::chrono::seconds>(duration).count();
        auto nanos = std::chrono::duration_cast<std::chrono::nanoseconds>(duration).count() % 1000000000;
        timestamp->set_seconds(seconds);
        timestamp->set_nanos(nanos);
        log.set_allocated_timestamp(timestamp);

        std::lock_guard<std::mutex> lock(mutex_);
        log_queue_.push(log);
        cond_var_.notify_one();
    }

    void SetServerStopping() {
        std::lock_guard<std::mutex> lock(mutex_);
        is_server_stopping_ = true;
        cond_var_.notify_all();
    }

private:
    std::queue<LogEntry> log_queue_;
    std::mutex mutex_;
    std::condition_variable cond_var_;
    bool is_server_stopping_ = false;
};

struct Handle {
    std::unique_ptr<Server> server;
    PLCRuntimeServiceImpl* service;  // 使用指针
    const char* server_address;
};

void* plc_runtime_grpc_server_init(const ServerConfig* config) {
    Handle* handle = new Handle();
    handle->service = new PLCRuntimeServiceImpl();  // 创建服务实例
    grpc::EnableDefaultHealthCheckService(true);
    grpc::ServerBuilder builder;
    builder.AddListeningPort(config->server_address, grpc::InsecureServerCredentials());
    builder.RegisterService(handle->service);  // 注册服务实例
    handle->server = builder.BuildAndStart();
    handle->server_address = config->server_address;
    return handle;
}

void plc_runtime_grpc_server_start(void* handle) {
    Handle* h = static_cast<Handle*>(handle);
    std::cout << "Server listening on " << h->server_address << std::endl;
    h->server->Wait();
}

void plc_runtime_grpc_server_stop(void* handle) {
    Handle* h = static_cast<Handle*>(handle);
    h->service->SetServerStopping();
    h->server->Shutdown();
}

void plc_runtime_grpc_server_free(void* handle) {
    Handle* h = static_cast<Handle*>(handle);
    delete h->service;  // 删除服务实例
    delete h;
}

// C风格接口函数实现
void log_message_rpc(void* handle, const char* component, const char* message, int level) {
    Handle* h = static_cast<Handle*>(handle);
    h->service->AddLog(component, message, static_cast<LogLevel>(level));
}
