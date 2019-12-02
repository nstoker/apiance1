package utils

import (
	"log"
	"os"
)

func init() {
	// if err := godotenv.Load(GetProjectRoot() + "/"); err != nil {
	// 	log.Printf(".env file not found in %s", GetProjectRoot())
	// }
}

// Getenv returns the environment variable
func Getenv(envVar string) string {
	value := os.Getenv(envVar)

	return value
}

// GetEnvWithError gets an enviroment variable, and throws an error if it is not not found
func GetEnvWithError(envVar string) string {
	value := Getenv(envVar)
	if value == "" {
		log.Fatalf("Environment variable '%s' missing", envVar)
	}

	return value
}
