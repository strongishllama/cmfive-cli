package cmd

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

var newModuleCmd = &cobra.Command{
	Use:     "module <module-name> [flags]",
	Short:   "Create a new module",
	Example: "cmfive new module payroll",
	Args:    checkArgs(1, "cannot create module: module name required"),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmfive.NewModule(args[0]); err != nil {
			return xerror.New("failed to create module", err)
		}

		return nil
	},
	SilenceErrors: true,
	// SilenceUsage:  true,
}

func init() {
	newCmd.AddCommand(newModuleCmd)

	f := func(cmd *cobra.Command) error {
		if newModuleCmd.Long != "" {
			cmd.Printf("%s\n\n", newModuleCmd.Long)
		} else {
			cmd.Printf("%s\n\n", newModuleCmd.Short)
		}

		cmd.Println("USAGE:")
		cmd.Printf("  %s\n\n", newModuleCmd.Use)

		flagUsages := newModuleCmd.LocalFlags().FlagUsages()
		if flagUsages != "" {
			cmd.Println("FLAGS:")
			cmd.Println(flagUsages)
		}

		cmd.Println("EXAMPLE:")
		cmd.Printf("  %s\n", newModuleCmd.Example)

		return nil
	}

	newModuleCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		f(cmd)
	})
	newModuleCmd.SetUsageFunc(f)
}
