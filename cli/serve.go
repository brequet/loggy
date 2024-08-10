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
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the Loggy server",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.OpenSQLiteDBIfExists("loggy.db")
			if err != nil {
				return fmt.Errorf("failed to open database: %v", err)
			}

			server := server.NewServer(config.SERVER_PORT, db, slog.Default())
			return server.Start()
		},
	}

	return cmd
}
