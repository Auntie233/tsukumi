package models

import (
	"time"
)

type UserInfo struct {
	CreateTime time.Time `xorm:"not null DATETIME"`
	Id         int       `xorm:"not null pk autoincr INT"`
	Mobile     string    `xorm:"VARCHAR(100)"`
	Password   string    `xorm:"VARCHAR(100)"`
	UpdateTime time.Time `xorm:"DATETIME"`
	Username   string    `xorm:"VARCHAR(100)"`
}
