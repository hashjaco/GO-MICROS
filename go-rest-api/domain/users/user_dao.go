package users

import (
	"fmt"
	"github.com/hashjaco/GO-MICROS/go-rest-api/database/mysql/users_db"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/errors"
	"github.com/hashjaco/GO-MICROS/go-rest-api/utils/mysql_utils"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "No rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, username, password, created_on, updated_on) VALUES(?, ?, ?, ?, ?, ?, ?);"
	queryGetUser     = "select id, first_name, last_name, email, username, password, created_on, updated_on FROM users WHERE id=?;"
	queryGetAllUsers = "select id, first_name, last_name, email, username, password, created_on, updated_on FROM users;"
)

/**
* Return User specified by ID number
**/
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.CreatedOn, &user.UpdatedOn); getErr != nil {
		return errors.NewNotFoundError(
			fmt.Sprintf("error when trying to get user %d: %s", user.ID, getErr.Error()))
	}
	return nil
}

/**
*  Return all Users and their respective rows from users table
**/
func (users *Users) GetAll() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetAllUsers)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	user := User{}

	// loop through addresses of users returned from database and append column values to collection
	for rows.Next() {
		var (
			id int64
			first_name, last_name, email, username, password, created_on, updated_on string
		)
		if err := rows.Scan(&id, &first_name, &last_name, &email, &username, &password, &created_on, &updated_on); err != nil {
			fmt.Printf("Gathering rows return error: %s" ,err.Error())

			return mysql_utils.ParseError(err)
		}
		user.ID = id
		user.FirstName = first_name
		user.LastName = last_name
		user.Email = email
		user.Username = username
		user.Password = password
		user.CreatedOn = created_on
		user.UpdatedOn = updated_on
		fmt.Println(user)

		users.collection = append(users.collection, user)
	}
	return nil
}

/**
* Insert User into database
**/
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Username, user.Password, user.CreatedOn, user.UpdatedOn)
	if saveErr != nil {
		if strings.Contains(saveErr.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("Email address %s already exists. Terminating transaction", user.Email))
		}
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while attrmting to save user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}
