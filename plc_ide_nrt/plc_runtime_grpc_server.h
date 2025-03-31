// plc_runtime_grpc_server_c.h
#ifndef PLC_RUNTIME_GRPC_SERVER_C_H
#define PLC_RUNTIME_GRPC_SERVER_C_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif


#define LOG_LEVEL_DEBUG 0
#define LOG_LEVEL_INFO 1
#define LOG_LEVEL_WARNING 2
#define LOG_LEVEL_ERROR 3
#define LOG_LEVEL_UNKNOWN 4

typedef struct {
    const char* server_address;
} ServerConfig;

void* plc_runtime_grpc_server_init(const ServerConfig* config);
void plc_runtime_grpc_server_start(void* handle);
void plc_runtime_grpc_server_stop(void* handle);
void plc_runtime_grpc_server_free(void* handle);

void log_message_rpc(void* handle, const char* component, const char* message, int level);

#ifdef __cplusplus
}
#endif

#endif  // PLC_RUNTIME_GRPC_SERVER_C_H
