package login

import (
	"github.com/gin-gonic/gin"
	"log"
	"tsukumi/app/common/db"
	"tsukumi/app/common/response"
	models "tsukumi/app/entity/po"
)

func Login(c *gin.Context) {
	resp := response.Gin{Ctx: c}
	//dicList := make([]models.DicInfo, 0)
	engine := db.GetConnect()
	//err := engine.ID(4).Find(&dicList)
	dic := &models.DicInfo{}
	_, err := engine.ID(4).Get(dic)
	if err != nil {
		log.Println(err.Error())
		resp.Success(err.Error(), nil)
		return
	}
	resp.Success("", dic)
}
