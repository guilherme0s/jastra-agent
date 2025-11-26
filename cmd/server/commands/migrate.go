package commands

import (
	"log/slog"

	"github.com/guilherme0s/atlans/pkg/config"
	"github.com/guilherme0s/atlans/pkg/store/postgres"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	RunE:  migrateCmdFn,
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

func migrateCmdFn(command *cobra.Command, args []string) error {
	slog.Info("Connecting to the database to run migrations...")

	dbSettings := config.DatabaseSettings{}
	dbSettings.SetDefaults()

	db, err := postgres.New(dbSettings)
	if err != nil {
		return err
	}
	defer db.Close()

	slog.Info("Running migrations...")
	if err := db.MigrateUp(dbSettings); err != nil {
		return err
	}

	return nil
}
