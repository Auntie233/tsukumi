package login

import (
	"github.com/gin-gonic/gin"
	"tsukumi/app/common/response"
)

type LoginVo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	vo := LoginVo{Username: "123", Password: "321"}
	resp := response.Gin{Ctx: c}
	resp.Success("", vo)
}
