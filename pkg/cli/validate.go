package cli

import (
	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"
)

// CheckArgs checks the correct number of command line arguments have been given.
func CheckArgs(length int, message string) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) > length {
			return xerror.New("too many arguments")
		}

		if len(args) < length {
			return xerror.New(message)
		}

		return nil
	}
}
