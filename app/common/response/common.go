package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tsukumi/app/entity/errno"
	"tsukumi/app/entity/result"
)

type Gin struct {
	Ctx *gin.Context
}

func (g *Gin) Success(message string, res interface{}) {
	g.Ctx.JSON(http.StatusOK, result.Result{
		Code:    result.CodeSuccess,
		Message: message,
		Data:    res})
	g.Ctx.Abort()
}

func (g *Gin) Failed(err *errno.Error) {
	g.Ctx.JSON(http.StatusOK, err)
	g.Ctx.Abort()
}
