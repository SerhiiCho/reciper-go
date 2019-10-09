package db

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// Database struct
type Database struct {
	*gorm.DB
	config *Config
}

// NewDatabase opens database
func NewDatabase(c *Config) *Database {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.DBUser, c.DBPwd, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open("mysql", dataSource)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(model.User{}, model.Recipe{})

	return &Database{db, c}
}
