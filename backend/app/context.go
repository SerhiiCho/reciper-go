package app

import (
	"github.com/SerhiiCho/reciper/backend/db"
	"github.com/SerhiiCho/reciper/backend/model"
)

type Context struct {
	RemoteAddress string
	Database      *db.Database
	User          *model.User
}
