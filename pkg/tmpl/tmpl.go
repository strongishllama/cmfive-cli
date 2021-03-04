package tmpl

import (
	"embed"
	"os"
	"strings"
	"text/template"

	"github.com/gofor-little/xerror"
)

// NewFileFromTemplate creates a file at filePath using the template at templatePath while
// parsing v into the template.
func NewFileFromTemplate(fileSystem embed.FS, templatePath string, filePath string, v interface{}) error {
	// Initialize the template with some helper functions mapped.
	tmpl := template.New("template").Funcs(template.FuncMap{
		"Title": strings.Title,
	})

	// Read the template file data.
	data, err := fileSystem.ReadFile(templatePath)
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
