package models

import (
	"time"
)

type DicInfo struct {
	Attribute1 string    `xorm:"VARCHAR(100)"`
	Attribute2 string    `xorm:"VARCHAR(100)"`
	Attribute3 string    `xorm:"VARCHAR(100)"`
	Code       string    `xorm:"not null unique(dic_info_UN) VARCHAR(100)"`
	CreateTime time.Time `xorm:"not null DATETIME"`
	DelFlag    int       `xorm:"not null default b'0' unique(dic_info_UN) BIT(1)"`
	Id         int       `xorm:"not null pk autoincr INT"`
	Name       string    `xorm:"not null VARCHAR(100)"`
	PCode      string    `xorm:"VARCHAR(100)"`
	Type       string    `xorm:"not null unique(dic_info_UN) VARCHAR(100)"`
	UpdateTime time.Time `xorm:"DATETIME"`
}
