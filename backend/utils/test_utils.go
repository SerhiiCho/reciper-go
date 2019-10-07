package utils

import "os"

// TestSetup sets env variables for testing database
func TestSetup() {
	loadEnvVariables()
}

func loadEnvVariables() {
	vars := map[string]string{
		"DB_HOST":     "localhost",
		"DB_USER":     "root",
		"DB_PASSWORD": "111111",
		"DB_PORT":     "1001",
		"DB_NAME":     "reciper",
	}

	for name, value := range vars {
		err := os.Setenv(name, value)
		HandleError("error setting env variable", err, "")
	}
}
