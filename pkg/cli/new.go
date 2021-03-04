package cli

import (
	"github.com/spf13/cobra"
)

// newCmd initializes and returns the command for CLI calls to "cmfive new".
func newCmd() *cobra.Command {
	command := &cobra.Command{
		Use:     "new <command> [flags]",
		Short:   "Create new Cmfive resources from templates",
		Example: "cmfive new module payroll",
	}

	command.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		if command.Long != "" {
			cmd.Printf("%s\n\n", command.Long)
		} else {
			cmd.Printf("%s\n\n", command.Short)
		}

		cmd.Println("USAGE:")
		cmd.Printf("  %s\n\n", command.Use)

		commands := command.Commands()
		if len(commands) != 0 {
			cmd.Println("COMMANDS:")

			for _, c := range commands {
				cmd.Printf("  %s\t\t%s\n", c.Name(), c.Short)
			}
		}
		cmd.Println()

		flagUsages := command.LocalFlags().FlagUsages()
		if flagUsages != "" {
			cmd.Println("FLAGS:")
			cmd.Println(flagUsages)
		}

		cmd.Println("EXAMPLE:")
		cmd.Printf("  %s\n", command.Example)
	})

	command.AddCommand(newModuleCmd())
	command.AddCommand(newActionCmd())
	command.AddCommand(newMigrationCmd())
	command.AddCommand(newModelCmd())

	return command
}
