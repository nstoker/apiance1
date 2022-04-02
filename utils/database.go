package utils

import (
	"fmt"
	"log"
)

// GetDatabaseConnectionString gets the database string given
// the passed environment variables.
// The connection parameter
func GetDatabaseConnectionString(connection, host, name, pass, port, user string) string {
	if connection == "" {
		connection = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
	}
	log.Printf("connection: %s", connection)
	return connection
}

// GetDatabaseConnection returns the database connection string
func GetDatabaseConnection() string {
	conn := Getenv("DATABASE_URL") // Set by Heroku
	host := Getenv("DATABASE_HOST")
	name := Getenv("DATABASE_NAME")
	pass := Getenv("DATABASE_PASSWORD")
	port := Getenv("DATABASE_PORT")
	user := Getenv("DATABASE_USER")

	return GetDatabaseConnectionString(conn, host, name, pass, port, user)
}
