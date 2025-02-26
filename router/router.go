package router

import (
	"mall/controller"
	"mall/middlewwares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

	trade := r.Group("/trade")
	{
		trade.GET("/list", controller.ShowItems)
		trade.GET("/search", controller.SearchItem)
		trade.Use(middlewwares.AuthMiddlewares()) // 需要登录
		trade.POST("/create", controller.CreateItem)
	}

	return r
}
