package app

import (
	"github.com/hashjaco/GO-MICROS/go-rest-api/controllers/users"
)

func mapUrls() {
	router.GET("/users", users.GetUsers)
	router.GET("/users/:user_id", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
