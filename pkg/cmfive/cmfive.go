package cmfive

import "os"

var (
	// TemplatesDir is the directory that the templates are located.
	// By default this is an empty string but can be overridden.
	TemplatesDir string = ""
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
