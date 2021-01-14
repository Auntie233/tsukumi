package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				panic(fmt.Sprintf("%s", r))
			}
		}()
		context.Next()
	}
}
