package main

import (
	"log"

	_ "github.com/arif14377/koda-b6-backend1/cmd/docs"
	"github.com/arif14377/koda-b6-backend1/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Testing Swagger API
// @version		1.0.0
// @description testing implementasi swagger
// @host		localhost:8888
// @BasePath	/
// @accept		json
// @produce		json

func main() {

	r := gin.Default()

	// AUTH
	rg := r.Group("/")
	{
		users := rg.Group("/users")
		{
			// register
			users.POST("register", handler.Register)
			// login
			users.POST("login", handler.Login)
			// get all users
			users.GET("/", handler.GetUsers)
			// check user details
			users.GET(":id", handler.UserDetails)
			// delete user
			users.DELETE(":id", handler.DeleteUser)
			// update data user
			users.PUT("/profile", handler.UpdateUser)
		}
		products := rg.Group("/products")
		{
			// PRODUCT
			// get all products
			products.GET("/", handler.GetProducts)
			// get product details
			products.GET(":id", handler.ProductDetails)
			// add product
			products.POST("/", handler.AddProduct)
			// delete product
			products.DELETE(":id", handler.DeleteProduct)
			products.PUT(":id", handler.UpdateProduct)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run("localhost:8888"); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
