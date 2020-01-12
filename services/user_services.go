package services

import (
	"github.com/kabbali/bookstore_users_api/domain/users"
	"github.com/kabbali/bookstore_users_api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User)  (*users.User, *errors.RestErr){
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User)  (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
		currentUser.FirstName = user.FirstName
	}


	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

func DeleteUser(userId int64) *errors.RestErr  {
	user := &users.User{Id: userId}
	return user.Delete()
}