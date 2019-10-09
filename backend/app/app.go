package app

import (
	"github.com/SerhiiCho/reciper/backend/db"
	"github.com/SerhiiCho/reciper/backend/utils"
)

type App struct {
	Database *db.Database
}

func NewApp(database *db.Database) (app *App) {
	app = &App{
		Database: database,
	}

	return app
}

func (app *App) Close() {
	err := app.Database.Close()
	utils.HandleError("", err, "")
}
