package db

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, name)

	db, err := gorm.Open("mysql", dataSource)
	utils.HandleError("Database connection error", err, "")

	return &Database{db}
}
