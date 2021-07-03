package cli

import (
	"bytes"
	"embed"
	"io"
	"os"
	"os/exec"

	"github.com/gofor-little/xerror"
)

var (
	//go:embed templates
	templates embed.FS
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cmd := rootCmd()
	cmd.AddCommand(newCmd())
	cmd.AddCommand(runCmd())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// execCommand will execute a command returning stderr and the error on failure
// or stdout and nil on success.
func execCommand(name string, args ...string) (string, error) {
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	// Use multi-writer so we can see the output in real time as well as capture it.
	mwStdout := io.MultiWriter(os.Stdout, &stdout)
	mwStderr := io.MultiWriter(os.Stderr, &stderr)

	cmd := exec.Command(name, args...)
	cmd.Stdout = mwStdout
	cmd.Stderr = mwStderr

	if err := cmd.Run(); err != nil {
		return stderr.String(), xerror.Wrap("failed to run command", err)
	}

	return stdout.String(), nil
}
