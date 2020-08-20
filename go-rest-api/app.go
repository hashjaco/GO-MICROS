package go_rest_api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) InitRouter() {
	app.Router.Methods("GET").Path("/users").HandlerFunc(app.getUsers)
	app.Router.Methods("POST").Path("/users").HandlerFunc(app.addUser)
}


func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {

	// Preparing query string with each argument to avoid any sort of SQL injection
	stmt, err := app.Database.Prepare("SELECT id, first_name, last_name, age, is_admin, created_on, updated_on FROM users")

	// check for any errors returned from preparing query
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	// manually close users if necessary after error is returned
	defer rows.Close()
	if err := json.NewEncoder(w).Encode(rows); err != nil {
		log.Fatal("An error occured while writing response to output")
	}
}

func (app *App) addUser(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("INSERT INTO demoDB (first_name, last_name, age, is_admin) VALUES ('hashim', 'jacobs', 33, true)")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	w.WriteHeader(http.StatusOK)
}
