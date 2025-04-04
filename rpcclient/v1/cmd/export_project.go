package cmd

import (
	"fmt"
	"log"

	"github.com/nihilKnight/grpc-section/client"
	plcruntime "github.com/nihilKnight/grpc-section/gen"
	"github.com/spf13/cobra"
)

func NewExportProjectCmd() *cobra.Command {
	var (
		versionID     uint32
		versionString string
		chunkSize     uint32
	)

	cmd := &cobra.Command{
		Use:   "export-project",
		Short: "Export PLC project",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn := client.GetConnection()
			defer conn.Close()

			client := plcruntime.NewPLCRuntimeServiceClient(conn)
			resp, err := client.ExportProject(cmd.Context(), &plcruntime.ProjectExportRequest{
				VersionId:     versionID,
				VersionString: versionString,
				ChunkSize:     chunkSize,
			})

			if err != nil {
				log.Fatalf("Error exporting project: %v", err)
			}
			fmt.Printf("Exported Project Metadata: %+v\n", resp)
			return nil
		},
	}

	cmd.Flags().Uint32Var(&versionID, "version-id", 0, "Version ID to export")
	cmd.Flags().StringVar(&versionString, "version-string", "", "Version string identifier")
	cmd.Flags().Uint32VarP(&chunkSize, "chunk-size", "s", 0, "Chunk size for export")

	cmd.MarkFlagRequired("version-id")
	cmd.MarkFlagRequired("version-string")
	cmd.MarkFlagRequired("chunk-size")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewExportProjectCmd())
}