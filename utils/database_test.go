package utils

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestGetDatabaseConnection(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		t.Fatalf("Can't load `.env.test`")
	}

	got := GetDatabaseConnection()
	want := "postgres://postgres:changeme@localhost:5432/apiance_test?sslmode=disable"
	if got != want {
		t.Errorf("Incorrect response: got %s, want %s", got, want)
	}
}
