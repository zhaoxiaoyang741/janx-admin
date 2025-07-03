package model

type User struct {
	BaseColumnWithStatus
	Username     string  `json:"username" gorm:"column:username;type:char(40);unique;not null;comment:'用户名'"`
	Password     string  `json:"password" gorm:"column:password;type:char(255);not null; comment:'密码'"`
	Nickname     string  `json:"nickname" gorm:"column:nickname;type:char(50);comment:'昵称'"`
	Gender       uint    `json:"gender" gorm:"column:gender;type:smallint;comment:'性别 0:保密 1:男 2:女'"`
	Avatar       string  `json:"avatar" gorm:"column:avatar;type:char(255);comment:'头像'"`
	Email        string  `json:"email" gorm:"column:email;type:char(100);comment:'邮箱'"`
	Phone        string  `json:"phone" gorm:"column:phone;type:char(11);unique;not null;comment:'手机号'"`
	Introduction string  `json:"introduction" gorm:"column:introduction;type:char(255);comment:'个人简介'"`
	Roles        []*Role `gorm:"many2many:user_roles" json:"roles"`
}
