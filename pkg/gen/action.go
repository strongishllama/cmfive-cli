package gen

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofor-little/xerror"

	"github.com/strongishllama/cmfive-cli/pkg/tmpl"
)

// actionData holds the data required to build the action template.
type actionData struct {
	Name       string
	Method     string
	ModuleName string
}

// NewAction creates a new action for a module.
func NewAction(moduleName string, actionName string, actionMethod string) error {
	moduleName = strings.ToLower(moduleName)
	actionName = strings.ToLower(actionName)
	actionMethod = strings.ToUpper(actionMethod)
	actionsDir := fmt.Sprintf("modules/%s/actions", moduleName)
	actionPath := fmt.Sprintf("%s/%s.php", actionsDir, actionName)

	// Check that the module directory exists.
	dirExists, err := exists(fmt.Sprintf("modules/%s", moduleName))
	if err != nil {
		return xerror.Wrap("failed to check if directory exists", err)
	}
	if !dirExists {
		return xerror.Wrap("module with that name doesn't exist", err)
	}

	// Create actions directory if it doesn't exist.
	dirExists, err = exists(actionsDir)
	if err != nil {
		return xerror.Wrap("failed to check if directory exists", err)
	}
	if !dirExists {
		if err := os.Mkdir(actionsDir, os.ModePerm); err != nil {
			return xerror.Wrap("failed to create directory", err)
		}
	}

	// Check that an action with that name doesn't already exist.
	actionExists, err := exists(actionPath)
	if err != nil {
		return xerror.Wrap("failed to check if file exists", err)
	}
	if actionExists {
		return xerror.Wrap("an action with that name already exists", err)
	}

	data := &actionData{Name: actionName, Method: actionMethod, ModuleName: moduleName}
	if err := tmpl.NewFileFromTemplate(templates, "templates/action.tmpl", actionPath, data); err != nil {
		return xerror.Wrap("failed to create file from template", err)
	}

	return nil
}
