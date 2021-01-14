package router

import (
	"github.com/gin-gonic/gin"
	"tsukumi/app/controller/login"
)

func InitRouter(r *gin.Engine) {
	r.POST("/login", login.Login)
}
