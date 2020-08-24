package users

import (
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
	"strings"
)

type User struct {
	ID int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

type Users struct {
	collection []User
}


func (user *User)Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	if user.Password == ""{
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}