package utils

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestGetDatabaseConnection(t *testing.T) {
	if err := godotenv.Load("../test.env"); err != nil {
		t.Fatalf("Can't load `test.env`")
	}

	got := GetDatabaseConnection()
	want := "postgres://apiance:changeme@127.0.0.1:5432/apiance1_api_test?sslmode=disable"
	if got != want {
		t.Errorf("Incorrect response: got %s, want %s", got, want)
	}
}
