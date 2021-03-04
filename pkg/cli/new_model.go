package cli

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

// newModelCmd initializes and returns the command for CLI calls to "cmfive new model".
func newModelCmd() *cobra.Command {
	command := &cobra.Command{
		Use:     "model <module-name> <model-name> [flags]",
		Short:   "Create a new model for a module",
		Example: "cmfive new model payroll Employee",
		Args:    CheckArgs(2, "cannot create action: module name and model name required"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := gen.NewModel(args[0], args[1]); err != nil {
				return xerror.New("failed to create model", err)
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
