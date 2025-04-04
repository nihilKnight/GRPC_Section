package cmd

import (
	"fmt"
	"log"

	"github.com/nihilKnight/grpc-section/client"
	plcruntime "github.com/nihilKnight/grpc-section/gen"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewCreateProjectCmd() *cobra.Command {
	var (
		projectID      uint32
		name           string
		description    string
		version        string
		taskNum        uint32
		tasksJSON      []string
		deviceNum      uint32
		devicesJSON    []string
	)

	cmd := &cobra.Command{
		Use:   "create-project",
		Short: "Create a new PLC project",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 解析Tasks
			var tasks []*plcruntime.TaskMetadata
			for _, t := range tasksJSON {
				task := &plcruntime.TaskMetadata{}
				if err := protojson.Unmarshal([]byte(t), task); err != nil {
					return fmt.Errorf("failed to parse task JSON: %v", err)
				}
				tasks = append(tasks, task)
			}

			// 验证Task数量
			if len(tasks) != int(taskNum) {
				return fmt.Errorf("task_num %d does not match provided tasks count %d", taskNum, len(tasks))
			}

			// 解析Devices
			var devices []*plcruntime.BusDeviceMetadata
			for _, d := range devicesJSON {
				device := &plcruntime.BusDeviceMetadata{}
				if err := protojson.Unmarshal([]byte(d), device); err != nil {
					return fmt.Errorf("failed to parse device JSON: %v", err)
				}
				devices = append(devices, device)
			}

			// 验证Device数量
			if len(devices) != int(deviceNum) {
				return fmt.Errorf("device_num %d does not match provided devices count %d", deviceNum, len(devices))
			}

			// 构建请求
			req := &plcruntime.ProjectMetadata{
				ProjectId:     projectID,
				Name:          name,
				Description:   description,
				VersionString: version,
				TaskNum:       taskNum,
				Tasks:         tasks,
				DeviceNum:     deviceNum,
				Devices:       devices,
			}

			// 调用gRPC
			conn := client.GetConnection()
			defer conn.Close()
			resp, err := plcruntime.NewPLCRuntimeServiceClient(conn).CreateProject(cmd.Context(), req)
			if err != nil {
				log.Fatalf("Error creating project: %v", err)
			}
			fmt.Printf("Created Project: %+v\n", resp)
			return nil
		},
	}

	cmd.Flags().Uint32Var(&projectID, "project-id", 0, "Project ID")
	cmd.Flags().StringVar(&name, "name", "", "Project name")
	cmd.Flags().StringVar(&description, "description", "", "Project description")
	cmd.Flags().StringVar(&version, "version", "v1.0.0", "Project version")
	cmd.Flags().Uint32Var(&taskNum, "task-num", 0, "Number of tasks")
	cmd.Flags().StringSliceVar(&tasksJSON, "task", []string{}, `Task metadata in JSON format (e.g. '{"task_id":1,"name":"main","type":"RISING"}')`)
	cmd.Flags().Uint32Var(&deviceNum, "device-num", 0, "Number of devices")
	cmd.Flags().StringSliceVar(&devicesJSON, "device", []string{}, `Device metadata in JSON format (e.g. '{"device_id":1,"type":"GPIO"}')`)

	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("task-num")
	cmd.MarkFlagRequired("device-num")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewCreateProjectCmd())
}