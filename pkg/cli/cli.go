package cli

import "os"

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cmd := rootCmd()
	cmd.AddCommand(newCmd())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
