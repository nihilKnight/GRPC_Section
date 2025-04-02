cmake_minimum_required(VERSION 3.10)
project(GRPC_Section)

# 设置第三方库路径
set(THIRD_PARTY_DIR ${CMAKE_SOURCE_DIR}/third_party)

# 配置 Protobuf 和 gRPC
find_package(Protobuf REQUIRED 
    HINTS ${THIRD_PARTY_DIR}/protobuf/cmake
)

find_package(gRPC REQUIRED
    HINTS ${THIRD_PARTY_DIR}/grpc/cmake
)

# 设置公共包含目录
include_directories(
    ${Protobuf_INCLUDE_DIRS}
    ${gRPC_INCLUDE_DIRS}
    ${THIRD_PARTY_DIR}/protobuf/include
    ${THIRD_PARTY_DIR}/grpc/include
)

# 添加子目录
add_subdirectory(plc_ide_nrt)
