package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/nstoker/apiance1/api/controllers"
	"github.com/nstoker/apiance1/api/migrate"
	"github.com/nstoker/apiance1/api/models"
	"github.com/nstoker/apiance1/utils"
)

var server = controllers.Server{}
var userInstance = models.User{}

// var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	log.Printf("$PWD '%s'", os.ExpandEnv("${PWD}"))
	err = godotenv.Load("../../test.env")
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	if err = server.InitializeDatabase(utils.GetDatabaseConnection()); err != nil {
		log.Printf("Error initializing database: %s", err)
		os.Exit(1)
	}

	if err := dropTables(); err != nil {
		log.Printf("Error dropping tables: %s", err)
		os.Exit(2)
	}

	if err := migrate.Perform(); err != nil {
		log.Printf("Error migrating tables: %s", err)
		os.Exit(3)
	}

	if err = server.InitializeRouter(); err != nil {
		log.Printf("Error initialising router: %s", err)
		os.Exit(4)
	}

	// We (probably) don't need to run a seeder on the test database.

	os.Exit(m.Run())
}

func dropTables() error {
	tables := []string{"users"}

	for _, t := range tables {
		log.Printf("Dropping %s", t)
		_, err := server.DB.Exec("DROP TABLE  $1", t)
		if err != nil {
			return fmt.Errorf("Error dropping '%s': %w", t, err)
		}
	}

	return nil
}

func refreshUserTable() error {
	// err := server.DB.DropTableIfExists(&models.User{}).Error
	// if err != nil {
	// 	return err
	// }
	// err = server.DB.AutoMigrate(&models.User{}).Error
	// if err != nil {
	// 	return err
	// }

	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	user := models.User{
		Name:     "Pet",
		Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
		Password: "password",
	}

	// err := server.DB.Model(&models.User{}).Create(&user).Error
	// if err != nil {
	// 	log.Fatalf("cannot seed users table: %v", err)
	// }
	return user, fmt.Errorf("seedOneUser Not Implemented")
}

func seedUsers() ([]models.User, error) {

	users := []models.User{
		models.User{
			Name:     "Steven victor",
			Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
			Password: "password",
		},
		models.User{
			Name:     "Kenny Morris",
			Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
			Password: "password",
		},
	}

	err := fmt.Errorf("seedUses Not Implemented")
	// for i := range users {
	// 	err := server.DB.Model(&models.User{}).Create(&users[i]).Error
	// 	if err != nil {
	// 		return []models.User{}, err
	// 	}
	// }
	return users, err
}
