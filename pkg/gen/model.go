package gen

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofor-little/xerror"

	"github.com/strongishllama/cmfive-cli/pkg/tmpl"
)

// modelData holds the data required to build the model template.
type modelData struct {
	Name string
}

// NewModel creates a new model for a module.
func NewModel(moduleName string, modelName string) error {
	moduleName = strings.ToLower(moduleName)
	modelName = fmt.Sprintf("%s%s", strings.Title(moduleName), modelName)
	modelsDir := fmt.Sprintf("modules/%s/models", moduleName)
	modelPath := fmt.Sprintf("%s/%s.php", modelsDir, modelName)

	// Check that the module directory exists.
	dirExists, err := exists(fmt.Sprintf("modules/%s", moduleName))
	if err != nil {
		return xerror.New("failed to check if directory exists", err)
	}
	if !dirExists {
		return xerror.New("module with that name doesn't exist", err)
	}

	// Create models directory if it doesn't exist.
	dirExists, err = exists(modelsDir)
	if err != nil {
		return xerror.New("failed to check if directory exists", err)
	}
	if !dirExists {
		if err := os.Mkdir(modelsDir, os.ModePerm); err != nil {
			return xerror.New("failed to create directory", err)
		}
	}

	// Check that an model with that name doesn't already exist.
	model, err := exists(modelPath)
	if err != nil {
		return xerror.New("failed to check if file exists", err)
	}
	if model {
		return xerror.New("a model with that name already exists", err)
	}

	if err := tmpl.NewFileFromTemplate(templates, "templates/model.tmpl", modelPath, &modelData{Name: modelName}); err != nil {
		return xerror.New("failed to create file from template", err)
	}

	return nil
}
