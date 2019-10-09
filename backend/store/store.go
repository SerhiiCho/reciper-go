package store

import (
	"database/sql"
	"fmt"
	"github.com/SerhiiCho/reciper/backend/utils"
	_ "github.com/go-sql-driver/mysql"
)

// Store struct
type Store struct {
	config *Config
	conn   *sql.DB
}

// NewStore opens database
func NewStore(conf *Config) *Store {
	user := conf.DBUser
	pwd := conf.DBPwd
	host := conf.DBHost
	port := conf.DBPort
	name := conf.DBName

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, name)
	db, err := sql.Open("mysql", dataSource)

	utils.FatalIfError(err)

	return &Store{conf, db}
}

func (store *Store) Open() {
	//
}

func (store *Store) Close() {
	//
}
