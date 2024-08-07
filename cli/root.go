package cli

import (
	"log/slog"
	"os"

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
	rootCmd.AddCommand(newServeCommand())

	return rootCmd
}
