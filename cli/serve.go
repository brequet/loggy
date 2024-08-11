package cli

import (
	"fmt"
	"log/slog"

	"github.com/brequet/loggy/config"
	"github.com/brequet/loggy/database"
	"github.com/brequet/loggy/server"
	"github.com/spf13/cobra"
)

func newServeCommand() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the Loggy server",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.OpenSQLiteDBIfExists("loggy.db")
			if err != nil {
				return fmt.Errorf("failed to open database: %v", err)
			}

			conf, err := config.LoadConfig(configFile, slog.Default())
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			server := server.NewServer(conf.Server.Port, db, slog.Default())
			return server.Start()
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to the configuration file")

	return cmd
}
