package gen

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofor-little/xerror"

	"github.com/strongishllama/cmfive-cli/pkg/tmpl"
)

// migrationData holds the data required to build the migration template.
type migrationData struct {
	Name string
}

// NewMigration creates a new migration for a module.
func NewMigration(moduleName string, migrationName string) error {
	moduleName = strings.ToLower(moduleName)
	migrationsDir := fmt.Sprintf("modules/%s/install/migrations", moduleName)

	// Check that the module directory exists.
	dirExists, err := exists(fmt.Sprintf("modules/%s", moduleName))
	if err != nil {
		return xerror.Wrap("failed to check if directory exists", err)
	}
	if !dirExists {
		return xerror.Wrap("module with that name doesn't exist", err)
	}

	// Create migrations directory if it doesn't exist.
	dirExists, err = exists(migrationsDir)
	if err != nil {
		return xerror.Wrap("failed to check if directory exists", err)
	}
	if !dirExists {
		if err := os.Mkdir(migrationsDir, os.ModePerm); err != nil {
			return xerror.Wrap("failed to create directory", err)
		}
	}

	migrationPath := fmt.Sprintf("%s/%s-%s%s.php", migrationsDir, time.Now().Format("20060102150405"), strings.Title(moduleName), migrationName)
	data := &migrationData{Name: fmt.Sprintf("%s%s", strings.Title(moduleName), migrationName)}
	if err := tmpl.NewFileFromTemplate(templates, "templates/migration.tmpl", migrationPath, data); err != nil {
		return xerror.Wrap("failed to create file from template", err)
	}

	return nil
}
