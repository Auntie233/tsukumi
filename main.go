package main

import (
	"github.com/gin-gonic/gin"
	"tsukumi/app/config"
	"tsukumi/app/middleware"
	"tsukumi/app/router"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(middleware.LoggerToFile(), middleware.Recover())
	router.InitRouter(engine)
	if err := engine.Run(config.PORT); err != nil {
		panic("服务启动失败")
	}
}
