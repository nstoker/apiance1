package api

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nstoker/apiance1/api/controllers"
	"github.com/nstoker/apiance1/api/migrate"
	"github.com/nstoker/apiance1/api/seed"
	"github.com/nstoker/apiance1/utils"
	"github.com/sirupsen/logrus"
)

var server = controllers.Server{}

// Run server, run.
func Run() error {
	err := godotenv.Load()
	if err != nil {
		logrus.Info(err)
	}

	if err := server.InitializeDatabase(utils.GetDatabaseConnection()); err != nil {
		return err
	}

	if err := server.InitializeRouter(); err != nil {
		return err
	}

	if err := migrate.Up(); err != nil {
		return err
	}

	if err := seed.Load(server.DB); err != nil {
		return err
	}
	port := os.Getenv("PORT")
	return server.Run(":" + port)
}
