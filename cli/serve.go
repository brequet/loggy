package cli

import (
	"log/slog"

	"github.com/brequet/loggy/server"
	"github.com/spf13/cobra"
)

func newServeCommand() *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the Loggy server",
		RunE: func(cmd *cobra.Command, args []string) error {
			srv, err := server.NewServer(port, slog.Default())
			if err != nil {
				return err
			}
			return srv.Start()
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on") // TODO: do not specify port

	return cmd
}
