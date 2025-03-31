cmake_minimum_required(VERSION 3.10)
project(runtime_main)

# 设置 C++ 标准
set(CMAKE_CXX_STANDARD 11)

# 查找 gRPC
find_package(gRPC REQUIRED)

# 查找 Protobuf
find_package(Protobuf REQUIRED)

# 包含头文件路径
include_directories(
    ${CMAKE_CURRENT_SOURCE_DIR}
    ${gRPC_INCLUDE_DIRS}
    ${Protobuf_INCLUDE_DIRS}
)

# gRPC 接口和 Protobuf 实现文件
set (UTIL_SOURCES
    plc_ide_nrt_v1.grpc.pb.cc
    plc_ide_nrt_v1.pb.cc
)

# 服务端源文件
set(SEVER_SOURCES
    test_grpc_server.c
    plc_runtime_grpc_server.cc
)

# 客户端源文件
set(CLIENT_SOURCES
    test_grpc_client.cc
)

# 添加可执行文件
add_executable(runtime_main ${SERVER_SOURCES} ${UTIL_SOURCES})
add_executable(ide_main ${CLIENT_SOURCES} ${UTIL_SOURCES})

# 链接 gRPC 和 Protobuf 库
target_link_libraries(runtime_main
    grpc++
    grpc
    protobuf
    pthread
    dl # For dynamic linking if needed
)

target_link_libraries(runtime_main
    grpc++
    grpc
    protobuf
    pthread
    dl # For dynamic linking if needed
)

# 可选：如果需要动态链接 gRPC 库
# set_target_properties(runtime_main PROPERTIES LINKER_LANGUAGE CXX) # 确保链接器使用 C++ 链接
