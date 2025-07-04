package model

type Role struct {
	BaseColumnWithStatus
	Name    string  `gorm:"type:char(20);not null;unique" json:"name"`
	Keyword string  `gorm:"type:char(20);not null;unique" json:"keyword"`
	Desc    *string `gorm:"type:char(100);" json:"desc"`
	Status  uint    `gorm:"type:smallint;default:1;comment:'1正常, 2禁用'" json:"status"`
	Sort    uint    `gorm:"type:smallint;default:999;comment:'角色排序(排序越大权限越低, 不能查看比自己序号小的角色, 不能编辑同序号用户权限, 排序为1表示超级管理员)'" json:"sort"`
	Users   []*User `gorm:"many2many:user_roles" json:"users"`
	// Menus   []*Menu `gorm:"many2many:role_menus;" json:"menus"` // 角色菜单多对多关系
}
