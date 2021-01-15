package models

import (
	"time"
)

type OauthApprovals struct {
	Clientid       string    `xorm:"VARCHAR(255)"`
	Expiresat      time.Time `xorm:"TIMESTAMP"`
	Lastmodifiedat time.Time `xorm:"TIMESTAMP"`
	Scope          string    `xorm:"VARCHAR(255)"`
	Status         string    `xorm:"VARCHAR(10)"`
	Userid         string    `xorm:"VARCHAR(255)"`
}
