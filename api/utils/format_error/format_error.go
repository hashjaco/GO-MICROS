package format_error

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("username already taken")
	}
	if strings.Contains(err, "email") {
		return errors.New("this email address is already registered")
	}
	return errors.New("incorrect details")
}
