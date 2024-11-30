package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs database migrations, if needed.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runMigrate(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println("Migration successful")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func runMigrate() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error while trying to get current working directory: %w", err)
	}

	db, err := getDb(dir)
	if err != nil {
		return fmt.Errorf("Error while trying to get current working directory: %w", err)
	}

	ctx := context.Background()
	goose.RunContext(ctx, "up", db, dir)

	return nil
}

func getDb(dir string) (*sql.DB, error) {
	// TODO: try to load .env file from current directory and create a DB connection
	return nil, fmt.Errorf("Not implemented")
}
