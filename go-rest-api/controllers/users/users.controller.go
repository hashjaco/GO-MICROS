package users

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context){
	c.String(http.StatusOK, "Users retrieved")
}
func SearchUser(c *gin.Context){
	c.String(http.StatusOK, "Searched for User")
}
func CreateUser(c *gin.Context){
	c.String(http.StatusOK, "Created New User")
}


