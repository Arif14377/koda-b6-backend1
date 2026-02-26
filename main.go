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
	FullName string `json:"fullName"`
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

	r.POST("/register", func(ctx *gin.Context) {
		data := Users{}
		err := ctx.ShouldBindJSON(&data)
		isExist := false

		if err != nil {
			ctx.JSON(401, Response{
				Success: false,
				Message: "JSON tidak valid.",
			})
			return
		}

		// validasi email:
		// 1. email harus ada @
		// 2. fullname, email dan password tidak boleh kosong
		// 3. Jika email sudah terdaftar, maka tidak bisa register
		// 4. Selain itu register berhasil.

		if !strings.Contains(data.Email, "@") {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Email tidak valid.",
			})
			return
		}

		if data.FullName == "" || data.Email == "" || data.Password == "" {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Data tidak boleh kosong.",
			})
			return
		}

		for _, u := range listUsers {
			if data.Email == u.Email {
				isExist = true
			}
		}

		if isExist {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Email sudah terdaftar.",
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

	r.POST("/login", func(ctx *gin.Context) {
		var data Users
		err := ctx.ShouldBindJSON(&data)
		login := false

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "JSON tidak valid",
			})
			return
		}

		if !strings.Contains(data.Email, "@") {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Email tidak valid.",
			})
			return
		}

		if data.Email == "" || data.Password == "" {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Email dan Password tidak boleh kosong.",
			})
			return
		}

		for _, u := range listUsers {
			if u.Email == data.Email {
				if u.Email == data.Email && u.Password == data.Password {
					login = true
				} else {
					ctx.JSON(400, Response{
						Success: false,
						Message: "Password salah.",
					})
					return
				}
			}
		}

		if login {
			ctx.JSON(200, Response{
				Success: true,
				Message: fmt.Sprintf("Welcome %s", data.Email),
			})
		} else {
			ctx.JSON(401, Response{
				Success: false,
				Message: "Email tidak terdaftar. Silahkan register terlebih dahulu.",
			})
		}

	})

	// TODO: buat handle update data user

	r.Run("localhost:8888")
}
