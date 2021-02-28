package cmfive

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/gofor-little/xerror"
)

var (
	moduleDirs = []string{
		"actions",
		"assets",
		"install/migrations",
		"models",
		"partials/actions",
		"partials/templates",
		"templates",
		"tests/acceptance",
		"tests/unit",
	}
	//go:embed templates
	templates embed.FS
)

// NewModule creates a new module in the modules directory. If the modules
// directory doesn't exist, one will be created.
func NewModule(name string) error {
	name = strings.ToLower(name)
	moduleDir := fmt.Sprintf("modules/%s", name)

	// Create modules directory if it doesn't exist.
	dirExists, err := exists("modules")
	if err != nil {
		return xerror.New("failed to check if directory exists", err)
	}
	if !dirExists {
		if err := os.Mkdir("modules", os.ModePerm); err != nil {
			return xerror.New("failed to create directory", err)
		}
	}

	// Check that that module name is free.
	dirExists, err = exists(moduleDir)
	if err != nil {
		return xerror.New("failed to check if directory exists", err)
	}
	if dirExists {
		return xerror.New("a module with that name already exists, module names must be uneque", nil)
	}

	// Create the directories required for a module.
	for _, d := range moduleDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", moduleDir, d), os.ModePerm); err != nil {
			return xerror.New("failed to create directory", err)
		}
	}

	// Create the files required for a module.
	moduleFiles := map[string]string{
		"templates/action.tmpl":   fmt.Sprintf("modules/%s/actions/index.php", name),
		"templates/service.tmpl":  fmt.Sprintf("modules/%s/models/%sService.php", name, strings.Title(name)),
		"templates/template.tmpl": fmt.Sprintf("modules/%s/templates/index.tpl.php", name),
		"templates/config.tmpl":   fmt.Sprintf("modules/%s/%s.config.php", name, name),
		"templates/hooks.tmpl":    fmt.Sprintf("modules/%s/%s.hooks.php", name, name),
		"templates/roles.tmpl":    fmt.Sprintf("modules/%s/%s.roles.php", name, name),
		"templates/README.tmpl":   fmt.Sprintf("modules/%s/README.md", name),
	}
	for k, v := range moduleFiles {
		if err := newFileFromTemplate(k, v); err != nil {
			_ = os.RemoveAll(moduleDir)
			return xerror.New("failed to create file from template", err)
		}
	}

	// Initialize git for the module.
	cmd := exec.Command("git", "init")
	cmd.Dir = moduleDir
	if _, err := cmd.Output(); err != nil {
		return xerror.New("failed to initialise git for module", err)
	}

	return nil
}

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

// newFileFromTemplate creates a file at filePath using the template at templatePath.
func newFileFromTemplate(templatePath string, filePath string) error {
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
	type module struct {
		Name string
	}
	if err := tmpl.Execute(file, &module{Name: "test"}); err != nil {
		return xerror.New("failed to execute template", err)
	}

	return nil
}
