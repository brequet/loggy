package cli

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/brequet/loggy/database"
	"github.com/brequet/loggy/ingester"
	"github.com/brequet/loggy/parser"
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

			ingestService, err := initializeForIngester()
			if err != nil {
				return err
			}

			return ingestService.IngestLogs(inputDir)
		},
	}

	return cmd
}

func initializeForIngester() (*ingester.Ingester, error) {
	db, err := database.NewSQLiteDB("loggy.db")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	err = db.CleanLogEntries()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to clean log entries: %v", err)
	}

	parseService := parser.NewParser()
	ingestService := ingester.NewIngester(db, parseService, slog.Default())

	return ingestService, nil
}
