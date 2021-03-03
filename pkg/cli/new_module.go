package cli

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

// newModuleCmd initializes and returns the command for CLI calls to "cmfive new module".
func newModuleCmd() *cobra.Command {
	command := &cobra.Command{
		Use:     "module <module-name> [flags]",
		Short:   "Create a new module",
		Example: "cmfive new module payroll",
		Args:    CheckArgs(1, "cannot create module: module name required"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := gen.NewModule(args[0]); err != nil {
				return xerror.New("failed to create module", err)
			}

			return nil
		},
		SilenceErrors: true,
	}

	f := func(cmd *cobra.Command) error {
		if command.Long != "" {
			cmd.Printf("%s\n\n", command.Long)
		} else {
			cmd.Printf("%s\n\n", command.Short)
		}

		cmd.Println("USAGE:")
		cmd.Printf("  %s\n\n", command.Use)

		flagUsages := command.LocalFlags().FlagUsages()
		if flagUsages != "" {
			cmd.Println("FLAGS:")
			cmd.Println(flagUsages)
		}

		cmd.Println("EXAMPLE:")
		cmd.Printf("  %s\n", command.Example)

		return nil
	}

	command.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		f(cmd)
	})
	command.SetUsageFunc(f)

	return command
}
