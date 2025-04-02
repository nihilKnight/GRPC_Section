cmake_minimum_required(VERSION 3.10)
project(runtime_main)

find_package(absl REQUIRED)
find_package(Protobuf REQUIRED COMPONENTS libprotobuf libprotoc)
find_package(gRPC REQUIRED)

# 设置 C++ 标准
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set_source_files_properties(
    test_grpc_server.c
    PROPERTIES
    LANGUAGE C
)

include_directories(
    $(Protobuf_INCLUDE_DIRS)
    $(grpc_INCLUDE_DIRS)
    $(absl_INCLUDE_DIRS)
)

set(CMAKE_PREFIX_PATH /home/went/.local)

# 本地文件路径
set(PROTO_SRC_DIR ${CMAKE_CURRENT_SOURCE_DIR})

# 生成 proto 文件的规则（如果需要动态生成）
# protobuf_generate_cpp(PROTO_SRCS PROTO_HDRS
#     plc_ide_nrt_v1.proto
# )

# 源文件集合
set(UTIL_SOURCES
    ${PROTO_SRC_DIR}/plc_ide_nrt_v1.grpc.pb.cc
    ${PROTO_SRC_DIR}/plc_ide_nrt_v1.pb.cc
)

set(SERVER_SOURCES
    ${PROTO_SRC_DIR}/test_grpc_server.c
    ${PROTO_SRC_DIR}/plc_runtime_grpc_server.cc
)

set(CLIENT_SOURCES
    ${PROTO_SRC_DIR}/test_grpc_client.cc
)

# 可执行文件
add_executable(runtime_main ${SERVER_SOURCES} ${UTIL_SOURCES})
add_executable(ide_main ${CLIENT_SOURCES} ${UTIL_SOURCES})

# 链接 gRPC 和 Protobuf 库
target_link_libraries(runtime_main
    absl::any
    gRPC::grpc++
    gRPC::grpc
    protobuf::libprotobuf
    protobuf::libprotoc
    Threads::Threads
    ${CMAKE_DL_LIBS}
)

target_link_libraries(ide_main
    absl::any
    gRPC::grpc++
    gRPC::grpc
    protobuf::libprotobuf
    protobuf::libprotoc
    Threads::Threads
    ${CMAKE_DL_LIBS}
)