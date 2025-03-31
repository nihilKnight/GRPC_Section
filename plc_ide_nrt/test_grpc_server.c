#include "plc_runtime_grpc_server.h"
#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <stdlib.h>

// 线程参数结构体
typedef struct {
    int thread_id;
    void* handle;
    const char* component;
    const char* message;
    int level;
    pthread_mutex_t* mutex;
} ThreadArgs;

int total_threads = 3;
int completed;

// 日志线程函数
void* log_thread(void* arg) {
    ThreadArgs* args = (ThreadArgs*)arg;
    sleep(3+args->thread_id); // 模拟延时
    log_message_rpc(args->handle, args->component, args->message, args->level);

    pthread_mutex_lock(args->mutex);
    completed ++;
    if (completed == total_threads) {
        plc_runtime_grpc_server_stop(args->handle);
    }
    pthread_mutex_unlock(args->mutex);

    free(args);
    return NULL;
}

int main() {
    ServerConfig config = { "0.0.0.0:50051" };
    void* handle = plc_runtime_grpc_server_init(&config);

    pthread_t thread1, thread2, thread3;

    // 互斥锁
    pthread_mutex_t mutex;
    pthread_mutex_init(&mutex, NULL);

    // 创建线程参数
    ThreadArgs* args1 = (ThreadArgs*)malloc(sizeof(ThreadArgs));
    args1->handle = handle;
    args1->component = "example one";
    args1->message = "Status is normal.";
    args1->level = LOG_LEVEL_INFO;
    args1->mutex = &mutex;

    ThreadArgs* args2 = (ThreadArgs*)malloc(sizeof(ThreadArgs));
    args2->handle = handle;
    args2->component = "example two";
    args2->message = "Debug is ok.";
    args2->level = LOG_LEVEL_DEBUG;
    args2->mutex = &mutex;

    ThreadArgs* args3 = (ThreadArgs*)malloc(sizeof(ThreadArgs));
    args3->handle = handle;
    args3->component = "example three";
    args3->message = "Fatal error!";
    args3->level = LOG_LEVEL_ERROR;
    args3->mutex = &mutex;

    // 创建线程
    pthread_create(&thread1, NULL, log_thread, args1);
    pthread_create(&thread2, NULL, log_thread, args2);
    pthread_create(&thread3, NULL, log_thread, args3);

    plc_runtime_grpc_server_start(handle);

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);
    pthread_join(thread3, NULL);

    plc_runtime_grpc_server_free(handle);

    return 0;
}
