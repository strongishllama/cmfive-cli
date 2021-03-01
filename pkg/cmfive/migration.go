package cmfive

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofor-little/xerror"
)

// migrationData holds the data required to build migration templates.
type migrationData struct {
	Name string
}

// NewMigration creates a new migration for a module.
func NewMigration(moduleName string, migrationName string) error {
	moduleName = strings.ToLower(moduleName)
	moduleDir := fmt.Sprintf("modules/%s", moduleName)
	migrationDir := fmt.Sprintf("modules/%s/install/migrations", moduleName)

	// Check that the module directory exists.
	dirExists, err := exists(moduleDir)
	if err != nil {
		return xerror.New("failed to check if directory exists", err)
	}
	if !dirExists {
		return xerror.New("module with that name doesn't exist", err)
	}

	// Create migrations directory if it doesn't exist.
	dirExists, err = exists(migrationDir)
	if err != nil {
		return xerror.New("failed to check if directory exists", err)
	}
	if !dirExists {
		if err := os.Mkdir(migrationDir, os.ModePerm); err != nil {
			return xerror.New("failed to create directory", err)
		}
	}

	migrationPath := fmt.Sprintf("%s/%s-%s%s.php", migrationDir, time.Now().Format("20060102150405"), strings.Title(moduleName), migrationName)
	data := &migrationData{Name: fmt.Sprintf("%s%s", strings.Title(moduleName), migrationName)}
	if err := newFileFromTemplate("templates/migration.tmpl", migrationPath, data); err != nil {
		return xerror.New("failed to create file from template", err)
	}

	return nil
}
