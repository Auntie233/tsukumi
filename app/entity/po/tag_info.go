package models

type TagInfo struct {
	Id   int    `xorm:"not null pk autoincr INT"`
	Mode string `xorm:"comment('心情：dic_info') VARCHAR(100)"`
	Name string `xorm:"unique(tag_info_UN) VARCHAR(100)"`
	Type string `xorm:"unique(tag_info_UN) VARCHAR(100)"`
}
