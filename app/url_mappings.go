package app

import (
	"github.com/gusmanwidodo/gobookstore_users-api/controllers/ping"
	"github.com/gusmanwidodo/gobookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}