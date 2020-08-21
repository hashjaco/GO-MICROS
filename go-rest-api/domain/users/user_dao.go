package users

import (
	"fmt"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
)

var(
	usersDB = make(map[int64]*User)
)


func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Age = result.Age
	user.Email = result.Email
	user.Password = result.Password
	user.Username = result.Username
	user.IsAdmin = result.IsAdmin
	user.CreatedOn = result.CreatedOn
	user.UpdatedOn = result.UpdatedOn
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s has already been registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}