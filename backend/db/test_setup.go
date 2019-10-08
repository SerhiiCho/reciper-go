package db

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/utils"
	"os"
)

// TestSetup sets env variables for testing database
func TestSetup() {
	loadEnvVariables()
	refreshDatabase()
}

func refreshDatabase() {
	DB := GetDB()
	query := utils.FileGetContent("../../db/migrations/recipe-seed.sql")

	// Get new Transaction. See http://golang.org/pkg/database/sql/#DB.Begin
	context, err := DB.Begin()
	utils.HandleError("Error while trying to begin the transaction", err, "")

	defer rollbackTransaction(context)

	_, execErr1 := context.Exec("TRUNCATE TABLE recipes;")
	_, execErr2 := context.Exec(query)

	utils.HandleError("Can't seed recipes table", execErr1, "")
	utils.HandleError("Can't seed recipes table", execErr2, "")
}

// rollbackTransaction Rolls back the transaction after function return.
// If transaction was already committed, this will do nothing.
func rollbackTransaction(context *sql.Tx) {
	rollErr := context.Rollback()
	utils.HandleError("Rollback transaction error", rollErr, "")
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
		utils.HandleError("error setting env variable", err, "")
	}
}
