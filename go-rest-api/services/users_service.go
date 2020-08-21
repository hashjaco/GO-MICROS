package services

import (
	"fmt"
	"github.com/hashjaco/GO-MICROS/go-rest-api/domain/users"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
)

func GetUsers() {
}

func GetUser(userID int64) (*users.User, *errors.RestErr){
	if userID <= 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("%d is an invalid user ID", userID))
	}
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr){
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
