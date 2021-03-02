package cmd

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/cmfive"
)

// newModuleCmd represents the newModule command
var newModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Create a new module",
	Args:  checkArgs(1, "cannot create module: module name required"),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		if err := cmfive.NewModule(args[0]); err != nil {
			return xerror.New("failed to create module", err)
		}

		return nil
	},
}

func init() {
	newCmd.AddCommand(newModuleCmd)
}
