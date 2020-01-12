package app

import (
	"github.com/kabbali/bookstore_users_api/controllers/ping"
	"github.com/kabbali/bookstore_users_api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.Create)
}