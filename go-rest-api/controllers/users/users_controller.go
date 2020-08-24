package users

import (
	"github.com/gin-gonic/gin"
	"github.com/hashjaco/GO-MICROS/go-rest-api/domain/users"
	users2 "github.com/hashjaco/GO-MICROS/go-rest-api/services/users"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
	"net/http"
	"strconv"
)

/**
* Handle GET requests to return all users
**/
func GetAllUsers(c *gin.Context){
	var userTable users.Users

	allUsers, err := users2.GetAllUsers(userTable)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, allUsers)
}

/**
* Handle GET requests with ID parameter
**/
func GetUser(c *gin.Context){
	// parse integer64 value passed into endpoint as a parameter and store in userID variable
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	// return error status code and message if userErr is defined above
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
	}

	user, getErr := users2.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}

	c.JSON(http.StatusOK, user)
}

/**
* Handle POST requests
**/
func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := users2.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}


