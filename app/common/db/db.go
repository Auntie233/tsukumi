package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"tsukumi/app/config"
)

var db *xorm.Engine

func getConnect() *xorm.Engine {
	if db == nil {
		var err error
		db, err = xorm.NewEngine("mysql", config.DbUrl)
		if err != nil {
			panic(fmt.Sprintf("数据库连接失败: %s", err.Error()))
		}
	}
	return db
}
