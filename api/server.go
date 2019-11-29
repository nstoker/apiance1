package api

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nstoker/apiance1/api/controllers"
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

	if err := server.Initialize(utils.GetDatabaseConnection()); err != nil {
		return fmt.Errorf("server.Run(): %w", err)
	}

	seed.Load(server.DB)
	port := os.Getenv("PORT")
	return server.Run(":" + port)
}
