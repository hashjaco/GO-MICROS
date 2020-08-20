package models

import (
	"net/http"
	"time"
)

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	IsAdmin string `json:"is_admin"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

func (app *App) InitRouter(){
	app.Router.Methods("GET").Path("/users").HandlerFunc(app.getUsers)
}

func (app *App) getUsers(w http.ResponseWriter, r *http.Request){
	users := &User{}

}