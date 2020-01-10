package app

import (
	"github.com/kabbali/bookstore_users_api/controllers/ping"
	"github.com/kabbali/bookstore_users_api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}