package model

type User struct {
	BaseColumnWithStatus
	Username     string `json:"username" gorm:"column:username;type:varchar(40);unique;not null;comment:'用户名'"`
	Password     string `json:"password" gorm:"column:password;type:varchar(255);not null; comment:'密码'"`
	Nickname     string `json:"nickname" gorm:"column:nickname;type:varchar(50);comment:'昵称'"`
	Avatar       string `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:'头像'"`
	Email        string `json:"email" gorm:"column:email;type:varchar(100);comment:'邮箱'"`
	Phone        string `json:"phone" gorm:"column:phone;type:varchar(11);unique;not null;comment:'手机号'"`
	Introduction string `json:"introduction" gorm:"column:introduction;type:varchar(255);comment:'个人简介'"`
}
