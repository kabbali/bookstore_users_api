package users

import (
	"github.com/kabbali/bookstore_users_api/datasources/mysql/users_db"
	"github.com/kabbali/bookstore_users_api/utils/date_utils"
	"github.com/kabbali/bookstore_users_api/utils/errors"
	"github.com/kabbali/bookstore_users_api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGerUser 	= "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryDeleteUser = "DELETE from users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryGerUser)
	if nil != err {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
		//if strings.Contains(err.Error(), noRowsInResultSet) {
		//	return errors.NewBadRequestError(
		//		fmt.Sprintf("user: %d does not exists", user.Id))
		//}
		//return errors.NewInternalServerError(
		//	fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr  {
	stmt, err := users_db.UsersDB.Prepare(queryInsertUser)
	if nil != err {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
		//if strings.Contains(err.Error(), indexUniqueEmail) {
		//	return errors.NewBadRequestError(
		//		fmt.Sprintf("email: %s already exists", user.Email))
		//}
		//return errors.NewInternalServerError(
		//	fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(
		//	fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryUpdateUser)
	if nil != err {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryDeleteUser)
	if nil != err {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}