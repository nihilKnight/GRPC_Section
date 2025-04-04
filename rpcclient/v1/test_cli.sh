# ExportLogs
./client-cli export-logs --chunk-size 1024

# CreateProject
./client-cli create-project \
  --project-id 1 \
  --name "MyProject" \
  --task-num 2 \
  --task '{"task_id":1, "name":"Task1", "type":"RISING", "priority":1, "cycle":100}' \
  --task '{"task_id":2, "name":"Task2", "type":"FALLING", "priority":2, "cycle":200}' \
  --device-num 1 \
  --device '{"device_id":1, "type":"GPIO", "cycle":50}'

# ImportProgram
./client-cli import-program \
  --program-id 100 \
  --program-name "DemoProgram" \
  --content-file ./demo.st

# ExportProject
./client-cli export-project --version-id 1 --version-string "v1.0.0" --chunk-size 2048
