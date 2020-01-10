package users

import (
	"fmt"
	"github.com/kabbali/bookstore_users_api/utils/errors"
)

var (
	mockDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	userInDB := mockDB[user.Id]
	if userInDB == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id 			= userInDB.Id
	user.FirstName 		= userInDB.FirstName
	user.LastName 		= userInDB.LastName
	user.Email 			= userInDB.Email
	user.DateCreated 	= userInDB.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr  {
	userInDB := mockDB[user.Id]
	if nil != userInDB {
		if user.Email == userInDB.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	mockDB[user.Id] = user
	return nil
}