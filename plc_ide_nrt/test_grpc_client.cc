// test_grpc_client.cc
#include "plc_ide_nrt_v1.grpc.pb.h"
#include "plc_ide_nrt_v1.pb.h"
#include <grpcpp/grpcpp.h>
#include <iostream>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using plcide::LogEntry;
using plcide::LogExportRequest;

int main() {
    const char* server_adderss = "0.0.0.0:50051";
    std::unique_ptr<plcide::PLCRuntimeService::Stub> stub = 
        plcide::PLCRuntimeService::NewStub(grpc::CreateChannel(server_adderss, 
        grpc::InsecureChannelCredentials()));
    

    LogExportRequest protoRequest;
    protoRequest.set_format("CSV");
    protoRequest.set_chunk_size(1024);

    ClientContext ctx;
    std::unique_ptr<grpc::ClientReader<plcide::LogEntry>> reader(stub->ExportLogs(&ctx, protoRequest));
    
    LogEntry log;
    while (reader->Read(&log)) {
        std::cout << log.component() << ": " << log.message() << "[" << log.level() << "]" << std::endl;
    }
}
