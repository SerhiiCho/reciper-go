package db

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

// Database struct
type Database struct {
	*gorm.DB
}

// NewDatabase opens database
func NewDatabase() *Database {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, name)

	db, err := gorm.Open("mysql", dataSource)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(model.User{}, model.Recipe{})

	return &Database{db}
}
