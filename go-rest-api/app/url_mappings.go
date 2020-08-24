package app

import (
	"github.com/hashjaco/GO-MICROS/go-rest-api/controllers/users"
)

func mapUrls() {
	router.GET("/users", users.GetAllUsers)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
