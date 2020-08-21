package users

import (
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
	"strings"
)

type User struct {
	ID int64 `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	Email string `json:"email"`
	IsAdmin bool `json:"is_admin"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

func (user *User)Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}