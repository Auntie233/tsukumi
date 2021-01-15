package models

type OauthClientDetails struct {
	AccessTokenValidity   int    `xorm:"INT"`
	AdditionalInformation string `xorm:"VARCHAR(4096)"`
	Authorities           string `xorm:"VARCHAR(255)"`
	AuthorizedGrantTypes  string `xorm:"VARCHAR(255)"`
	Autoapprove           string `xorm:"VARCHAR(255)"`
	ClientId              string `xorm:"not null pk VARCHAR(255)"`
	ClientSecret          string `xorm:"VARCHAR(255)"`
	RefreshTokenValidity  int    `xorm:"INT"`
	ResourceIds           string `xorm:"VARCHAR(255)"`
	Scope                 string `xorm:"VARCHAR(255)"`
	WebServerRedirectUri  string `xorm:"VARCHAR(255)"`
}
