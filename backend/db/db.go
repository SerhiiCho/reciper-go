package db

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

// Database struct
type Database struct {
	*gorm.DB
}

// New opens database
func New() *Database {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, name)

	db, err := gorm.Open("mysql", dataSource)
	utils.HandleError("Database connection error", err, "")

	db.AutoMigrate(model.User{}, model.Recipe{})

	return &Database{db}
}
