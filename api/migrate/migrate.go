package migrate

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/nstoker/apiance1/utils"

	// Because these are needed example said to...
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var migrationDir = fmt.Sprintf("file://%s/api/migrate/files", utils.GetProjectRoot())

// Down will reverse all of the migrations
func Down() error {
	m, err := migrate.New(migrationDir, utils.GetDatabaseConnection())
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := m.Down(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

// Up will perform the migrations.
func Up() error {
	m, err := migrate.New(migrationDir, utils.GetDatabaseConnection())
	if err != nil {
		return fmt.Errorf("migrate.Perform() Error preparing migrations: %w", err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("migrate.Perform() Error running migrations: %w", err)
		}
		log.Printf("Migration status: %v", err)
	}
	return nil
}
