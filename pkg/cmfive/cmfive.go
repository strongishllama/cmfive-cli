package cmfive

import (
	"embed"
	"os"
	"strings"
	"text/template"

	"github.com/gofor-little/xerror"
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

// newFileFromTemplate creates a file at filePath using the template at templatePath while
// parsing v into the template.
func newFileFromTemplate(templatePath string, filePath string, v interface{}) error {
	// Initialize the template with some helper functions mapped.
	tmpl := template.New("template").Funcs(template.FuncMap{
		"Title":   strings.Title,
		"ToUpper": strings.ToUpper,
	})

	// Read the template file data.
	data, err := templates.ReadFile(templatePath)
	if err != nil {
		return xerror.New("failed to read file data", err)
	}

	// Parse the template file data into the template.
	tmpl, err = tmpl.Parse(string(data))
	if err != nil {
		return xerror.New("failed to parse templates", err)
	}

	// Create the file that the template execution will be written to.
	file, err := os.Create(filePath)
	if err != nil {
		return xerror.New("failed to open file", err)
	}

	// Execute the template.
	if err := tmpl.Execute(file, v); err != nil {
		return xerror.New("failed to execute template", err)
	}

	return nil
}
