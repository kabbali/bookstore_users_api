package users

import (
	"fmt"
	"github.com/kabbali/bookstore_users_api/utils/errors"
	"regexp"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) ValidateUser() *errors.RestErr {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	if len(user.Email) > 254 || !rxEmail.MatchString(user.Email) {
		errorMessage := fmt.Sprintf("%s is not a valid email address", user.Email)
		return errors.NewBadRequestError(errorMessage)

	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}

//func ValidateUser(user *User) *errors.RestErr {
//	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
//	if user.Email == "" {
//		return errors.NewBadRequestError("invalid email address")
//	}
//	return nil
//}