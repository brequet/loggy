package cli

import (
	"errors"

	"github.com/spf13/cobra"
)

func newIngestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest <input_directory>",
		Short: "Ingest log files into the database",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputDir := args[0]
			if inputDir == "" {
				return errors.New("input directory is required")
			}

			ingestService, err := initializeServices()
			if err != nil {
				return err
			}

			return ingestService.IngestLogs(inputDir)
		},
	}

	return cmd
}
