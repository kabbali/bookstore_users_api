package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_username 	= "mysql_users_username"
	mysql_users_password 	= "mysql_users_password"
	mysql_users_host 		= "mysql_users_host"
	mysql_users_schema 		= "mysql_users_schema"
)

var (
	UsersDB *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username ,password, host, schema,
	)
	var err error
	UsersDB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = UsersDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}