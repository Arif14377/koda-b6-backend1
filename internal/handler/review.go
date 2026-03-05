package handler

import (
	"github.com/arif14377/koda-b6-backend1/internal/entity"
	"github.com/gin-gonic/gin"
)

var listReviews []entity.Review

// TODO: get all reviews
func getReviews(c *gin.Context) {
	if len(listReviews) != 0 {
		c.JSON(200, entity.Response{
			Success: true,
		})
	}
}
