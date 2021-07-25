package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gin-api/dao"
	"gin-api/middlewares"

	sentrygin "github.com/getsentry/sentry-go/gin"
	_ "gin-api/docs"
	"gin-api/routes"
	"gin-api/system"
)

func init() {
	system.LoadConfiguration()
	system.Sentry()
}

// @title gin-api API
// @version 2.0
// @description 接口文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 651019063@qq.com

// @license.name menghuanwd
// @license.url https://www.gin-api.com/

// @host api.gin-api.com
// @BasePath /
func main() {
	dao.Connect()
	defer dao.Close()

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	router.Use(middlewares.Cors())
	router.Use(middlewares.ParamsAnalyze())
	router.Use(middlewares.ErrHandler())

	routes.InitRoutes(router)

	router.Run(":4000")
}
