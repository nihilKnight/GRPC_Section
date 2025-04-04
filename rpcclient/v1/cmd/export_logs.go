package cmd

import (
	"fmt"
	"log"

	"github.com/nihilKnight/grpc-section/client"
	plcruntime "github.com/nihilKnight/grpc-section/gen"
	"github.com/spf13/cobra"
)

func NewExportLogsCmd() *cobra.Command {
	var chunkSize uint32

	cmd := &cobra.Command{
		Use:   "export-logs",
		Short: "Export PLC runtime logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn := client.GetConnection() // 实现gRPC连接
			defer conn.Close()

			client := plcruntime.NewPLCRuntimeServiceClient(conn)
			resp, err := client.ExportLogs(cmd.Context(), &plcruntime.LogExportRequest{
				ChunkSize: chunkSize,
			})

			if err != nil {
				log.Fatalf("Error exporting logs: %v", err)
			}
			fmt.Printf("Log Entry: %+v\n", resp)
			return nil
		},
	}

	cmd.Flags().Uint32VarP(&chunkSize, "chunk-size", "s", 0, "Chunk size for log export")
	cmd.MarkFlagRequired("chunk-size")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewExportLogsCmd())
}