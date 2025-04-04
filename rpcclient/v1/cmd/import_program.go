package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/nihilKnight/grpc-section/client"
	plcruntime "github.com/nihilKnight/grpc-section/gen"
	"github.com/spf13/cobra"
)

func NewImportProgramCmd() *cobra.Command {
	var (
		programID   uint32
		programName string
		contentFile string
	)

	cmd := &cobra.Command{
		Use:   "import-program",
		Short: "Import a PLC program",
		RunE: func(cmd *cobra.Command, args []string) error {
			contentBytes, err := os.ReadFile(contentFile)
			if err != nil {
				return fmt.Errorf("error reading content file: %v", err)
			}

			req := &plcruntime.ProgramSource{
				ProgramId:   programID,
				ProgramName: programName,
				Content:     string(contentBytes),
			}

			conn := client.GetConnection()
			defer conn.Close()
			resp, err := plcruntime.NewPLCRuntimeServiceClient(conn).ImportProgram(cmd.Context(), req)
			if err != nil {
				log.Fatalf("Error importing program: %v", err)
			}
			fmt.Printf("Import Status: %+v\n", resp)
			return nil
		},
	}

	cmd.Flags().Uint32Var(&programID, "program-id", 0, "Program ID")
	cmd.Flags().StringVar(&programName, "program-name", "", "Program name")
	cmd.Flags().StringVar(&contentFile, "content-file", "", "Path to program source file")
	
	cmd.MarkFlagRequired("program-id")
	cmd.MarkFlagRequired("program-name")
	cmd.MarkFlagRequired("content-file")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewImportProgramCmd())
}