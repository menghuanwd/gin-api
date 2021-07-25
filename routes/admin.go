package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers/merchant/category"
)

func AdminRoutes(router *gin.Engine) {

	api := router.Group("/admin")
	categories := api.Group("/categories")
	{
		categories.GET("/", category.Index)
	}
}
