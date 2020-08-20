package application

import (
	"database/sql"
	"github.com/hashjaco/go-websockets-tut/application/models"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Database *sql.DB
}

func (app *App) InitRouter(){
	app.Router.Methods("GET").Path("/users").HandlerFunc(app.getUsers)
}

