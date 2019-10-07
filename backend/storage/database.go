package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/SerhiiCho/reciper/backend/utils"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, name)

	DB, err := sql.Open("mysql", dataSource)
	utils.HandleError("Database connection error", err, "")

	return DB
}
