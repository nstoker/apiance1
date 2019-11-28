package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/neil-stoker/apiance1/api/controllers"
	"github.com/neil-stoker/apiance1/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}

// var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

// Database a database
func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	user := models.User{
		Name:     "Pet",
		Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
		Password: "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
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
	log.Printf("seedUsers: users: %+v", users)

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Printf("Error seeding %+v", i)
			return []models.User{}, err
		}
	}
	return users, nil
}

func refreshUserAndPostTable() error {

	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() error {

	err := refreshUserAndPostTable()
	if err != nil {
		return err
	}
	user := models.User{
		Name:     "Sam Phil",
		Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
		Password: "password",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return err
	}

	// err = server.DB.Model(&models.Post{}).Create(&post).Error
	// if err != nil {
	// 	return models.Post{}, err
	// }
	return nil
}

func seedUsersAndPosts() ([]models.User, error) {

	var err error

	if err != nil {
		return []models.User{}, err
	}
	var users = []models.User{
		models.User{
			Name:     "Steven victor",
			Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
			Password: "password",
		},
		models.User{
			Name:     "Magu Frank",
			Email:    fmt.Sprintf("%s@example.com", models.GenKsuid()),
			Password: "password",
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
	return users, nil
}
