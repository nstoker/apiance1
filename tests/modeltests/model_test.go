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
	err = godotenv.Load(fmt.Sprintf("%s/test.env", utils.GetProjectRoot()))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	dbName := os.Getenv("DATABASE_NAME")
	expectedDbName := "apiance1_api_test"
	if dbName != expectedDbName {
		log.Fatalf("Error getting database name, got '%s' want '%s'", dbName, expectedDbName)
	}

	if err = server.InitializeDatabase(utils.GetDatabaseConnection()); err != nil {
		log.Printf("Error initializing database: %s", err)
		os.Exit(1)
	}

	if err := migrate.Down(); err != nil {
		log.Printf("Error reversing migrations: %s", err)
		os.Exit(2)
	}

	if err := migrate.Up(); err != nil {
		log.Printf("Error migrating tables: %v", err)
		os.Exit(3)
	}

	if err = server.InitializeRouter(); err != nil {
		log.Printf("Error initialising router: %s", err)
		os.Exit(4)
	}

	// We (probably) don't need to run a seeder on the test database.

	os.Exit(m.Run())
}

func clearTables() error {
	// tables := []string{"users"}

	// for _, t := range tables {
	// 	log.Printf("Clearing '%s'", t)
	// 	sqlStatement := fmt.Sprintf("DELETE FROM %s;", t)
	// 	_, err := server.DB.Exec(sqlStatement)
	// 	if err != nil {
	// 		return fmt.Errorf("Error clearing '%s' with '%s': %w", t, sqlStatement, err)
	// 	}
	// }

	return nil
}

func refreshTable(table string) error {
	sqlStatement := fmt.Sprintf(`DELETE FROM %s;`, table)
	_, err := server.DB.Exec(sqlStatement)
	if err != nil {
		return fmt.Errorf("Error clearing %s: %v", table, err)
	}

	return nil
}

func seedOneUser() (*models.User, error) {

	refreshTable("users")

	user := models.User{
		Name:     "Pet",
		Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
		Password: "password",
	}

	newUser, err := user.CreateUser(server.DB)

	// err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("cannot seed users table: %v", err)
	}
	return newUser, nil
}

func seedUsers() ([]models.User, error) {
	newUsers := []models.User{}

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

	for _, u := range users {
		user, err := u.CreateUser(server.DB)
		if err != nil {
			return []models.User{}, err
		}

		newUsers = append(newUsers, *user)
	}
	return users, nil
}
