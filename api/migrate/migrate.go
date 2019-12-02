package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/nstoker/apiance1/utils"

	// Because these are needed example said to...
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Perform will perform the migrations.
func Perform() error {
	var migrationDir = fmt.Sprintf("file://%s/api/migrate/files", utils.GetProjectRoot())

	m, err := migrate.New(migrationDir, utils.GetDatabaseConnection())
	if err != nil {
		return fmt.Errorf("migrate.Perform() Error preparing migrations: %w", err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("migrate.Perform() Error running migrations: %w", err)
		}
	}

	return nil
}
