package handler

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/arif14377/koda-b6-backend1/internal/entity"
	"github.com/gin-gonic/gin"
)

var listUsers []entity.Users

// Register User
//
// @Summary      Create new user
// @Description  Register new user
// @Tags         create-user
// @Accept       json
// @Produce      json
// @Param		 name body		string	true		"Name"
// @Param		 email body		string	true		"Email"
// @Param		 password body	string	true		"Password"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /users/register [post]
func Register(ctx *gin.Context) {
	data := entity.Users{}
	err := ctx.ShouldBindJSON(&data)
	isExist := false

	if err != nil {
		ctx.JSON(401, entity.Response{
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
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Email tidak valid.",
		})
		return
	}

	if data.FullName == "" || data.Email == "" || data.Password == "" {
		ctx.JSON(400, entity.Response{
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
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Email sudah terdaftar.",
		})
		return
	}

	data.Id = len(listUsers) + 1
	listUsers = append(listUsers, data)
	ctx.JSON(200, entity.Response{
		Success: true,
		Message: "Registrasi berhasil.",
	})
}

// Login User
//
// @Summary      login
// @Description  Login with email and password
// @Tags         login-user
// @Accept       json
// @Produce      json
// @Param		 email body		string	true		"Email"
// @Param		 password body	string	true		"Password"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /users/login [post]
func Login(ctx *gin.Context) {
	var data entity.Users
	err := ctx.ShouldBindJSON(&data)
	login := false

	if err != nil {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "JSON tidak valid",
		})
		return
	}

	if !strings.Contains(data.Email, "@") {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "Email tidak valid.",
		})
		return
	}

	if data.Email == "" || data.Password == "" {
		ctx.JSON(400, entity.Response{
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
				ctx.JSON(400, entity.Response{
					Success: false,
					Message: "Password salah.",
				})
				return
			}
		}
	}

	if login {
		ctx.JSON(200, entity.Response{
			Success: true,
			Message: fmt.Sprintf("Welcome %s", data.Email),
		})
	} else {
		ctx.JSON(401, entity.Response{
			Success: false,
			Message: "Email tidak terdaftar. Silahkan register terlebih dahulu.",
		})
	}
}

// Get User
//
// @Summary      get list all user
// @Description  Show all user with struct list
// @Tags         all-user
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /users [get]
func GetUsers(ctx *gin.Context) {
	ctx.JSON(200, entity.Response{
		Success: true,
		Message: "List Users:",
		Results: listUsers,
	})
}

// Get User Details
//
// @Summary      get the details user
// @Description  Show details user
// @Tags         details-user
// @Accept       json
// @Produce      json
// @Param		 id path		int	true		"User ID"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /users/{id} [get]
func UserDetails(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := entity.Users{}
	notFound := true

	for _, u := range listUsers {
		if u.Id == id {
			user = u
			notFound = false
			break
		}
	}

	if notFound {
		ctx.JSON(404, entity.Response{
			Success: false,
			Message: "User tidak ditemukan",
		})
		return
	}

	ctx.JSON(200, entity.Response{
		Success: true,
		Message: fmt.Sprintf("data user ID: %d", id),
		Results: user,
	})
}

// Delete User
//
// @Summary      Delete user
// @Description  Delete user with param id
// @Tags         delete-user
// @Accept       json
// @Produce      json
// @Param		 id path		int	true		"User ID"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
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
		ctx.JSON(404, entity.Response{
			Success: false,
			Message: "User tidak ditemukan",
		})
		return
	}

	ctx.JSON(200, entity.Response{
		Success: true,
		Message: fmt.Sprintf("Data dengan id %d berhasil dihapus", id),
	})
}

// Update User
//
// @Summary      Update user
// @Description  Update user with param id
// @Tags         Update-user
// @Accept       json
// @Produce      json
// @Param		 id path		int	true		"User ID"
// @Success      200  {object}  entity.Response
// @Failure      400  {object}  entity.Response
// @Failure      404  {object}  entity.Response
// @Failure      500  {object}  entity.Response
// @Router       /users/{id} [put]
func UpdateUser(ctx *gin.Context) {
	data := entity.Users{}
	err := ctx.ShouldBindJSON(&data)
	// Validasi update data:
	// 1. Email yang sudah terdaftar tidak bisa dipakai

	if err != nil {
		ctx.JSON(400, entity.Response{
			Success: false,
			Message: "JSON tidak valid.",
		})
	}

	for _, u := range listUsers {
		if data.Email == u.Email {
			ctx.JSON(400, entity.Response{
				Success: false,
				Message: "Email sudah terdaftar.",
			})
			return
		}
	}

	listUsers = append(listUsers, data)
	ctx.JSON(200, entity.Response{
		Success: true,
		Message: "Data berhasil diperbarui.",
	})
}

// ARGON
// 1. install go get -u ...
// 2. sebelum password di masukkan ke dalam slice, lakukan hashing
// argon := argon2.DefaultConfig()
// hash, err := argon.Hash([]byte(data.Password), nil)
// di Append :
// ListUser = append(ListUser, Users{
// Email: data.mail,
// Password: string(hash.Enconde())
// })

// Login
// pembandingan
// ok, err := argon2.VerifyEncoded([]byte(data.Password), []found.Password)
// validasi.
// Jika benar kembalikan token. sementara kembalikan user.

// CORS
// r.OPTIONS(....){
// 2 di bawah ini adalah setting yang dibuat untuk header. setelahnya digunakan untuk mendapatkan isian body
// ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
// ctx.Header("Access-Control-Allow-Header", "http://localhost:5173")
// ctx.GetHeader()
// } - ditulis di setiap endpoint (kecuali pakai middleware)
// *Dokumentasi di MDN

// MIDDLEWARE
// 1. buat middlewar (misal
