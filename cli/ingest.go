package cli

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/brequet/loggy/config"
	"github.com/brequet/loggy/database"
	"github.com/brequet/loggy/ingester"
	"github.com/brequet/loggy/parser"
	"github.com/spf13/cobra"
)

func newIngestCommand() *cobra.Command {
	var configFile, targetDir string

	cmd := &cobra.Command{
		Use:   "ingest",
		Short: "Ingest log files into the database",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.NewSQLiteDB("loggy.db")
			if err != nil {
				return fmt.Errorf("failed to initialize database: %v", err)
			}

			err = db.CleanLogEntries()
			if err != nil {
				db.Close()
				return fmt.Errorf("failed to clean log entries: %v", err)
			}

			conf, err := config.LoadConfig(configFile, slog.Default())
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			parseService, err := parser.NewParser(conf.Parser.Formats, slog.Default())
			if err != nil {
				return fmt.Errorf("failed to initialize parser: %v", err)
			}

			ingestService := ingester.NewIngester(db, parseService, slog.Default())

			if err != nil {
				return err
			}

			if targetDir != "" {
				return ingestService.IngestLogs(targetDir)
			} else if len(conf.Parser.AppLogDirs) > 0 {
				return ingestService.IngestLogsForAppLogDirs(conf.Parser.AppLogDirs)
			} else {
				return errors.New("no target directory or app log dirs specified")
			}
		},
	}

	cmd.Flags().StringVarP(&targetDir, "target", "d", "", "Path to the target directory")
	cmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to the configuration file")

	return cmd
}
