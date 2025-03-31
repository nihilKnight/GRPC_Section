## What runtime can provide 

### Log

```
    switch (level) {
        case LOG_LEVEL_DEBUG:
        case LOG_LEVEL_INFO:
            printf("[%s] %s: %s >>> %s\n", get_level_string(level), time_str, component, content);
            break;
        case LOG_LEVEL_WARNING:
            printf("\033[33m[%s] %s: %s >>> %s\033[0m\n", get_level_string(level), time_str, component, content);
            break;
        case LOG_LEVEL_ERROR:
            printf("\033[31m[%s] %s: %s >>> %s\033[0m\n", get_level_string(level), time_str, component, content);
            // 如果有 dlerror() 信息，则输出
            printf("\033[31m[%s] %s：dlerror >>> %s\033[0m\n", get_level_string(level), time_str, dlerror());
            break;
        default:
            break;
    }
```

log(received): level, time, component, content

### Config 

#### files

config(could be upgraded?), and plc.so(need to be preprocessed on the ide-side)


```
typedef struct Project_obj{
    char            project_name[100];
    char            project_version[20];
    char            project_description[200];

    void*           handle;                 // 工程动态库句柄

    int             task_count;
    TaskPtr*        tasks;                  // 任务列表，指针数组
    TaskScheduler*  scheduler;              // 工程调度器

    int             bus_count;   
    BusDevicePtr*   bus_devices;            // 总线设备列表，指针数组
    pthread_mutex_t update_input_buffer_mutex;  // 输入缓冲区更新锁
    pthread_mutex_t update_output_buffer_mutex; // 输出缓冲区更新锁
} Project, *ProjectPtr;
```

about bus device:

```
// 通用的 Bus 设备结构
typedef struct BusDevice_obj {
    char name[32];                 // 设备名称（如 "GPIO", "Modbus"）
    BusProtocolType type;          // 设备类型
    int id;                        // 设备 ID
    int period;                    // 设备轮询周期，单位 ms
    pthread_t thread;              // 运行线程
    bool init_done;                // 初始化状态
    bool running;                  // 运行状态
    void *context;                              // 设备的上下文
    // void (*init)(void *context);                // 初始化函数
    // void (*update_input)(void *context);        // 输入更新函数，用于从设备IO缓冲区复制数据到IEC缓冲区
    // void (*process)(void *context);             // 处理函数，用于处理设备IO缓冲区的数据
    // void (*update_output)(void *context);       // 输出更新函数，用于从IEC缓冲区复制数据到设备IO缓冲区
    // void (*shutdown)(void *context);            // 关闭函数
    BusDeviceInit           init;
    BusDeviceUpdateInput    update_input;
    BusDeviceProcess        process;
    BusDeviceUpdateOutput   update_output;
    BusDeviceShutdown       shutdown;
    GlueVars                glue_vars;
} BusDevice, *BusDevicePtr;
```



