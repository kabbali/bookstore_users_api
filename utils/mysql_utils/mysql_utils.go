package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/kabbali/bookstore_users_api/utils/errors"
	"strings"
)

const (
	noRowsInResultSet = "no rows in result set"
	ErrDupEntry        = 1062
)

func ParseError(err error) *errors.RestErr  {
	driverErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsInResultSet) {
			return errors.NewNotFoundError("No matching record for the given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}
	switch driverErr.Number {
	case ErrDupEntry:
		return errors.NewBadRequestError("Duplicate entry: invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}