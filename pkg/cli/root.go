package cli

import (
	"github.com/spf13/cobra"
)

// rootCmd initializes and returns the commend for CLI calls to "cmfive".
func rootCmd() *cobra.Command {
	command := &cobra.Command{
		Use:     "cmfive <command> <subcommand> [flags]",
		Short:   "A CLI tool to help create Cmfive applications.",
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

	return command
}
