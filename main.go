package main

import (
	"fmt"
	"slices"
	"strconv"
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

	// TODO: buat validasi http error jika user id tidak ditemukan
	// TODO: data masih mengambil dari index, bukan dari Id Users
	r.GET("users/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		user := Users{}
		notFound := true

		for _, u := range listUsers {
			if u.Id == id {
				user = u
				notFound = false
				break
			}
		}

		if notFound {
			ctx.JSON(404, Response{
				Success: false,
				Message: "User tidak ditemukan",
			})
			return
		}

		ctx.JSON(200, Response{
			Success: true,
			Message: fmt.Sprintf("data user ID: %d", id),
			Results: user,
		})

	})

	// TODO: belum validasi jika user id tidak ditemukan
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		notFound := true

		for i, u := range listUsers {
			if u.Id == id {
				listUsers = slices.Delete(listUsers, i, i+1)
				notFound = false
				break
			}
		}

		if notFound {
			ctx.JSON(404, Response{
				Success: false,
				Message: "User tidak ditemukan",
			})
			return
		}

		ctx.JSON(200, Response{
			Success: true,
			Message: fmt.Sprintf("Data dengan id %d berhasil dihapus", id),
		})
	})

	// TODO: buat handle login

	r.Run("localhost:8888")
}
