package db

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	utils.FatalIfError(err)

	db.AutoMigrate(model.User{}, model.Recipe{})

	return &Database{db, c}
}
