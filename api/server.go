package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nstoker/apiance1/api/controllers"
	"github.com/nstoker/apiance1/api/migrate"
	"github.com/nstoker/apiance1/api/seed"
	"github.com/nstoker/apiance1/utils"
)

var server = controllers.Server{}

// Run server, run.
func Run() error {
	var err error
	err = godotenv.Load()
	if err != nil {
		return fmt.Errorf("server.Run() Error getting env, not comming through %w", err)
	}

	if err := server.InitializeDatabase(utils.GetDatabaseConnection()); err != nil {
		return err
	}

	if err := server.InitializeRouter(); err != nil {
		return err
	}

	if err := migrate.Perform(); err != nil {
		return err
	}
	if server.DB == nil {
		log.Fatal("server.DB is nil")
	}
	if err := seed.Load(server.DB); err != nil {
		return err
	}
	port := os.Getenv("PORT")
	return server.Run(":" + port)
}
