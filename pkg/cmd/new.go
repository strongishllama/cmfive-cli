package cmd

import (
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new <command> [flags]",
	Short:   "Create new Cmfive resources from templates",
	Example: "cmfive new module payroll",
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		if newCmd.Long != "" {
			cmd.Printf("%s\n\n", newCmd.Long)
		} else {
			cmd.Printf("%s\n\n", newCmd.Short)
		}

		cmd.Println("USAGE:")
		cmd.Printf("  %s\n\n", newCmd.Use)

		commands := newCmd.Commands()
		if len(commands) != 0 {
			cmd.Println("COMMANDS:")

			for _, c := range commands {
				cmd.Printf("  %s\t\t%s\n", c.Name(), c.Short)
			}
		}
		cmd.Println()

		flagUsages := newCmd.LocalFlags().FlagUsages()
		if flagUsages != "" {
			cmd.Println("FLAGS:")
			cmd.Println(flagUsages)
		}

		cmd.Println("EXAMPLE:")
		cmd.Printf("  %s\n", newCmd.Example)
	})

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
