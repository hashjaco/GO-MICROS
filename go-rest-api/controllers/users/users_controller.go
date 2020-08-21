package users

import (
	"github.com/gin-gonic/gin"
	"github.com/hashjaco/GO-MICROS/go-rest-api/domain/users"
	"github.com/hashjaco/GO-MICROS/go-rest-api/services"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context){
	c.String(http.StatusOK, "Users retrieved")
}

func GetUser(c *gin.Context){
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user ID")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}


