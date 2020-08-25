package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/hashjaco/GO-MICROS/api/utils/errors"
	"strings"
)

const(
	errorNoRows = "No rows in result set"
)

func ParseError(err error) *errors.RestErr{
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("Error processing request")
}
