package commands

import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/guilherme0s/atlans/pkg/config"
	"github.com/guilherme0s/atlans/pkg/server"
	"github.com/guilherme0s/atlans/pkg/store/postgres"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:  "server",
	RunE: serverCmdFn,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func serverCmdFn(command *cobra.Command, args []string) error {
	db, err := postgres.New(config.DatabaseSettings{})
	if err != nil {
		return err
	}
	defer db.Close()

	srv := server.New(config.ServerSettings{})

	idleConnsClosed := make(chan struct{})
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		srv.Shutdown()
		close(idleConnsClosed)
	}()

	slog.Info("Starting server...")
	if err = srv.Start(); err != http.ErrServerClosed {
		slog.Error("failed to start server", "err", err)
		return err
	}

	<-idleConnsClosed
	slog.Info("Server stopped")
	return nil
}
