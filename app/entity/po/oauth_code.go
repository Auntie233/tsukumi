package models

type OauthCode struct {
	Authentication []byte `xorm:"MEDIUMBLOB"`
	Code           string `xorm:"VARCHAR(255)"`
}
