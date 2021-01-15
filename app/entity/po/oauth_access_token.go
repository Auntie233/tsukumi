package models

type OauthAccessToken struct {
	Authentication   []byte `xorm:"MEDIUMBLOB"`
	AuthenticationId string `xorm:"not null pk VARCHAR(255)"`
	ClientId         string `xorm:"VARCHAR(255)"`
	RefreshToken     string `xorm:"VARCHAR(255)"`
	Token            []byte `xorm:"MEDIUMBLOB"`
	TokenId          string `xorm:"VARCHAR(255)"`
	UserName         string `xorm:"VARCHAR(255)"`
}
