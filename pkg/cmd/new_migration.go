package cmd

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

var newMigrationCmd = &cobra.Command{
	Use:     "migration <module-name> <migration-name> [flags]",
	Short:   "Create a new migration for a module",
	Example: "cmfive new migration payroll InitialMigration",
	Args:    checkArgs(2, "cannot create migration: module name migration name required"),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		if err := cmfive.NewMigration(args[0], args[1]); err != nil {
			return xerror.New("failed to create migration", err)
		}

		return nil
	},
	SilenceErrors: true,
	// SilenceUsage:  true,
}

func init() {
	newCmd.AddCommand(newMigrationCmd)

	f := func(cmd *cobra.Command) error {
		if newMigrationCmd.Long != "" {
			cmd.Printf("%s\n\n", newMigrationCmd.Long)
		} else {
			cmd.Printf("%s\n\n", newMigrationCmd.Short)
		}

		cmd.Println("USAGE:")
		cmd.Printf("  %s\n\n", newMigrationCmd.Use)

		flagUsages := newMigrationCmd.LocalFlags().FlagUsages()
		if flagUsages != "" {
			cmd.Println("FLAGS:")
			cmd.Println(flagUsages)
		}

		cmd.Println("EXAMPLE:")
		cmd.Printf("  %s\n", newMigrationCmd.Example)

		return nil
	}
	newMigrationCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		f(cmd)
	})
	newMigrationCmd.SetUsageFunc(f)
}
