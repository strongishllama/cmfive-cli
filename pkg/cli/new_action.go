package cli

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/gen"
)

// newActionCmd initializes and returns the command for CLI calls to "cmfive new action".
func newActionCmd() *cobra.Command {
	command := &cobra.Command{
		Use:     "action <module-name> <action-name> <action-method> [flags]",
		Short:   "Create a new action for a module",
		Example: "cmfive new action payroll edit GET",
		Args:    CheckArgs(3, "cannot create action: module name, action name and action method required"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := gen.NewAction(args[0], args[1], args[2]); err != nil {
				return xerror.New("failed to create action", err)
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
