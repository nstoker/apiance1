package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/nstoker/apiance1/api/utils"

	// Because these are needed example said to...
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Perform will perform the migrations.
func Perform() err {
	const migrationDir = "file://api/migrate/files"

	m, err := migrate.New(migrationDir, utils.GetDatabaseConnection())
	if err != nil {
		return fmt.Printf("migrate.Perform() Error preparing migrations: %w", err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Printf("migrate.Perform() Error running migrations: %w", err)
		}
	}

	return nil
}
