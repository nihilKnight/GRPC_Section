#!/bin/bash

# 测试配置
HOST="localhost:50051"
PROTO_FILE="proto/plcruntime.proto"
OUTPUT_DIR="reports-cli"
REPORT_FORMAT="html"
INSECURE=true

# 服务列表
declare -A SERVICES=(
  ["ExportProject"]="plcruntime.PLCRuntimeService.ExportProject"
)

# 公共参数
COMMON_ARGS=(
  "--proto=${PROTO_FILE}"
  "--format=${REPORT_FORMAT}"
  "--skipTLS"
  "--connections=4"
  "--cpus=$(nproc)"
  "--concurrency=50"
  "--total=1000"
  "--timeout=30s"
)

if [ "$INSECURE" = true ]; then
  COMMON_ARGS+=("--insecure")
fi

# 创建报告目录
mkdir -p "${OUTPUT_DIR}"

run_test() {
  local service_name=$1
  local method=$2
  local data_dir=$3
  
  echo "===================================================================="
  echo "🚀 开始测试服务: ${service_name}"
  echo "🕒 开始时间: $(date +'%Y-%m-%d %H:%M:%S')"
  echo "📂 数据目录: ${data_dir}"
  
  # 遍历测试数据文件
  for data_file in "${data_dir}"/*.json; do
    if [ ! -f "${data_file}" ]; then
      continue
    fi

    local base_name=$(basename "${data_file}" .json)
    local timestamp=$(date +%Y%m%d_%H%M%S)
    local report_file="${OUTPUT_DIR}/${service_name}_${base_name}_${timestamp}.html"

    echo "----------------------------------------"
    echo "🔧 测试配置:"
    echo "  方法: ${method}"
    echo "  数据文件: ${data_file}"
    echo "  报告文件: ${report_file}"

    # 执行测试命令
    ghz \
      --call="${method}" \
      "${HOST}" \
      --data-file="${data_file}" \
      -o "${report_file}" \
      "${COMMON_ARGS[@]}"

    # 检查执行结果
    if [ $? -eq 0 ]; then
      echo "✅ 测试成功完成"
    else
      echo "❌ 测试执行失败"
    fi
  done

  echo "🕒 结束时间: $(date +'%Y-%m-%d %H:%M:%S')"
  echo "====================================================================\n"
}

# 执行所有服务测试
for service in "${!SERVICES[@]}"; do
  data_dir="testdata/${service,,}"
  method="${SERVICES[$service]}"
  
  if [ -d "${data_dir}" ]; then
    run_test "${service}" "${method}" "${data_dir}"
  else
    echo "⚠️ 警告: 未找到测试数据目录 ${data_dir}"
  fi
done

echo "🎉 所有测试执行完成！报告已保存至 ${OUTPUT_DIR} 目录"
