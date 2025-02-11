package main

import (
	"mall/config"
	"mall/router"
)

func main() {
	config.InitConfig()       //初始化配置
	r := router.SetupRouter() //初始化路由

	port := config.AppConfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(port) // 参数即是运行的端口
}
