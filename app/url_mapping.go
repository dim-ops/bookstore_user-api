package app

import (
	"github.com/DimProject/bookstore_users-api/controllers/ping"
	"github.com/DimProject/bookstore_users-api/controllers/users"
)

func mapUrls() {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:user_id", users.GetUser)
	r.POST("/users/search", users.SearchUser)
	r.POST("/users", users.CreateUser)
}
