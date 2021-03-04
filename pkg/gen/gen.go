package gen

import (
	"embed"
	"os"
)

var (
	// TemplatesDir is the directory that the templates are located.
	// By default this is an empty string but can be overridden.
	TemplatesDir string = ""
	//go:embed templates
	templates embed.FS
)

// exosts checks if the file or directory exists at the given path.
func exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
