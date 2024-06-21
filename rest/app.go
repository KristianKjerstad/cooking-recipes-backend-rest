package rest

import (
	"net/http"
	"rest/database"
)

type App struct {
	Router   http.Handler
	Database database.DB
}

func New() App {
	app := App{Database: *database.Connect()}
	app.loadRoutes()
	return app
}
