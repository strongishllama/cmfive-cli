package cmd

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

// newMigrationCmd represents the newModule command
var newMigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Create a new migration for a module",
	Args:  checkArgs(2, "cannot create migration: module and migration name required"),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		if err := cmfive.NewMigration(args[0], args[1]); err != nil {
			return xerror.New("failed to create migration", err)
		}

		return nil
	},
}

func init() {
	newCmd.AddCommand(newMigrationCmd)
}
