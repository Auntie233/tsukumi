package models

type OauthClientToken struct {
	AuthenticationId string `xorm:"not null pk VARCHAR(255)"`
	ClientId         string `xorm:"VARCHAR(255)"`
	Token            []byte `xorm:"MEDIUMBLOB"`
	TokenId          string `xorm:"VARCHAR(255)"`
	UserName         string `xorm:"VARCHAR(255)"`
}
