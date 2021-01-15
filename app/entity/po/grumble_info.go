package models

import (
	"time"
)

type GrumbleInfo struct {
	Content     string    `xorm:"LONGTEXT" json:"content"`
	CreateTime  time.Time `xorm:"not null DATETIME" json:"createTime"`
	DelFlag     int       `xorm:"not null default b'0' BIT(1)" json:"delFlag"`
	Description string    `xorm:"VARCHAR(300)"`
	Id          int       `xorm:"not null pk autoincr INT"`
	Mode        string    `xorm:"comment('心情：dic') VARCHAR(100)"`
	Title       string    `xorm:"VARCHAR(100)"`
	Type        string    `xorm:"not null VARCHAR(20)"`
	UpdateTime  time.Time `xorm:"DATETIME"`
	UserId      int       `xorm:"not null INT"`
}
