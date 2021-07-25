package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers/merchant/category"
)

func MerchantRoutes(router *gin.Engine) {

	api := router.Group("/merchant")
	categories := api.Group("/categories")
	{
		categories.GET("/", category.Index)
		categories.POST("/", category.Create)
		categories.DELETE("/:id", category.Destroy)
		categories.GET("/:id", category.Show)
		categories.PUT("/:id", category.Update)
	}
}
