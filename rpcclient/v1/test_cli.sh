# ExportLogs
./client-cli export-logs --chunk-size 1024

# CreateProject
./client-cli create-project \
  --project-id 1 \
  --name "MyProject" \
  --task-num 0 \
  --device-num 0 \


# ImportProgram
./client-cli import-program \
  --program-id 100 \
  --program-name "DemoProgram" \
  --content-file ./demo/demo.st

# ExportProject
./client-cli export-project --version-id 1 --version-string "v1.0.0" --chunk-size 2048
