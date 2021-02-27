package cmd

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/cmfive"
)

var (
	// Name is the name of the module.
	Name string
)

// newModuleCmd represents the newModule command
var newModuleCmd = &cobra.Command{
	Use:   "module <name>",
	Short: "Create a new Cmfive module",
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

	// Here you will define your flags and configuration settings.
	// newModuleCmd.Flags().StringVar(&Name, "name")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newModuleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newModuleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkArgs(length int, message string) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) > length {
			return xerror.New("too many arguments", nil)
		}

		if len(args) < length {
			return xerror.New(message, nil)
		}

		return nil
	}
}
