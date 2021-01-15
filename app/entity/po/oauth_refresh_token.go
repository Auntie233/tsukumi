package models

type OauthRefreshToken struct {
	Authentication []byte `xorm:"MEDIUMBLOB"`
	Token          []byte `xorm:"MEDIUMBLOB"`
	TokenId        string `xorm:"VARCHAR(255)"`
}
