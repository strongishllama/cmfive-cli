package cli

import (
	"os"
	"strings"

	"github.com/gofor-little/xerror"
	"github.com/spf13/cobra"

	"github.com/strongishllama/cmfive-cli/pkg/tmpl"
)

// composerData holds the data required to build the composer template.
type composerData struct {
	Branch string
}

func runComposerCmd() *cobra.Command {
	command := &cobra.Command{
		Use:     "composer [flags]",
		Short:   "Run composer to install dependencies",
		Example: "cmfive run composer",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Figure out the installed PHP version.
			if err := tmpl.NewFileFromTemplate(templates, "templates/version.tmpl", "version.php", nil); err != nil {
				return xerror.Wrap("failed to create file from template", err)
			}
			defer os.Remove("version.php")

			output, err := execCommand("php", "version.php")
			if err != nil {
				return xerror.Wrap("failed to check PHP version", err)
			}

			parts := strings.Split(output, ".")
			if len(parts) != 2 {
				return xerror.New("unexptected length from parts")
			}

			// Build template to create composer.json file.
			data := composerData{Branch: "master"}
			if parts[0] == "7" && parts[1] == "0" {
				data.Branch = "legacy/PHP7.0"
			}

			if err := tmpl.NewFileFromTemplate(templates, "templates/composer.tmpl", "composer.json", data); err != nil {
				return xerror.Wrap("failed to create file from template", err)
			}

			// Install composer dependencies.
			output, err = execCommand("php", "composer.phar", "install")
			if err != nil {
				return xerror.Wrap("failed to install composer dependencies", err)
			}
			cmd.Println(output)

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
