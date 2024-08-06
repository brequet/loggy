package cli

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/brequet/loggy/database"
	"github.com/brequet/loggy/ingester"
	"github.com/brequet/loggy/parser"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var debug bool

	rootCmd := &cobra.Command{
		Use:   "loggy",
		Short: "Loggy is a tool for ingesting and analyzing log files",
		Long:  `A fast and flexible log ingestion and analysis tool built with Go.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logLevel := slog.LevelInfo
			if debug {
				logLevel = slog.LevelDebug
			}
			logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: logLevel,
			}))
			slog.SetDefault(logger)
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug mode")

	rootCmd.AddCommand(newIngestCommand())

	return rootCmd
}

func initializeServices() (*ingester.Ingester, error) {
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
