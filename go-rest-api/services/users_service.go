package services

import (
	"fmt"
	"github.com/hashjaco/GO-MICROS/go-rest-api/domain/users"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/date_utils"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
)

func GetAllUsers(userTable users.Users) (*users.Users, *errors.RestErr) {

	result := &users.Users{}
	if err := result.GetAll(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr){
	if userId <= 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("%d is an invalid user ID", userId))
	}
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr){
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.CreatedOn = date_utils.GetNowDBFormat()
	user.UpdatedOn = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
