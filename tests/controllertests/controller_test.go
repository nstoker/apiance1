package controllertests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/nstoker/apiance1/api/controllers"
	"github.com/nstoker/apiance1/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {
	log.Fatal("Database Not Implemented")
	// var err error

	// TestDbDriver := os.Getenv("TestDbDriver")

	// if TestDbDriver == "postgres" {
	// 	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
	// 	server.DB, err = sqlx.Open(TestDbDriver, DBURL)
	// 	if err != nil {
	// 		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
	// 		log.Fatal("This is the error:", err)
	// 	} else {
	// 		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	// 	}
	// }
}

func refreshTable(table string) error {
	err := fmt.Errorf("refreshUserTable not implemented")
	// server.DB.DropTableIfExists(&models.User{}).Error
	// if err != nil {
	// 	return err
	// }
	// err = server.DB.AutoMigrate(&models.User{}).Error
	// if err != nil {
	// 	return err
	// }

	return err
}

func seedOneUser() (*models.User, error) {

	err := refreshTable("users")
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Name:     "Pet",
		Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
		Password: "password",
	}

	newUser, err := user.CreateUser(server.DB)

	if err != nil {
		return &models.User{}, err
	}

	return newUser, nil
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
	for _, u := range users {
		user, err := u.CreateUser(server.DB)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, *user)
	}
	return users, nil
}
