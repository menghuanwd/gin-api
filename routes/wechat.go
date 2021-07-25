package routes

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers/wechat/category"
	"gin-api/controllers/wechat/user"
	"gin-api/middlewares"
)

func WechatRoutes(router *gin.Engine) {

	api := router.Group("/wechat")

	users := api.Group("/users")
	users.Use()
	{
		users.POST("/", user.Create)
		users.POST("/grant", user.Grant)
		users.POST("/wxacodeunlimit", user.CodeUnlimit)
	}
	users.Use(middlewares.Authenticate())
	{
		users.GET("/me", user.Me)
		users.PUT("/:id", user.Update)
	}

	categories := api.Group("/categories")
	{
		categories.GET("/", category.Index)
	}
}
