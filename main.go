package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

var listUsers []Users

func main() {
	r := gin.Default()

	r.POST("/users", func(ctx *gin.Context) {
		data := Users{}
		err := ctx.ShouldBindJSON(&data)

		if err != nil {
			ctx.JSON(401, Response{
				Success: false,
				Message: "You have to login first.",
			})
			return
		}

		for x := range listUsers {
			if strings.Contains(data.Email, listUsers[x].Email) {
				ctx.JSON(400, Response{
					Success: false,
					Message: "Email sudah terdaftar.",
				})
				return
			}
		}

		if data.Email == "" || data.Password == "" {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Email dan Password tidak boleh kosong.",
			})
			return
		}

		data.Id = len(listUsers) + 1
		listUsers = append(listUsers, data)
		ctx.JSON(200, Response{
			Success: true,
			Message: "Registrasi berhasil.",
		})

	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "List Users:",
			Results: listUsers,
		})
	})

	r.Run("localhost:8888")
}
