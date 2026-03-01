package main

import (
	"github.com/arif14377/koda-b6-backend1/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// register
	r.POST("/register", handler.Register)
	// login
	r.POST("/login", handler.Login)
	// get all users
	r.GET("/users", handler.GetUsers)
	// check user details
	r.GET("users/:id", handler.UserDetails)
	// delete user
	r.DELETE("/users/:id", handler.DeleteUser)
	// update data user
	r.PUT("/profile", handler.UpdateUser)

	r.Run("localhost:8888")
}
