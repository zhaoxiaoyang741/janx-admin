package service

import (
	"janx-admin/app/model"
	"janx-admin/app/vo"
	"janx-admin/global"
	"strings"
)

type RoleServiceInterface interface {
	Create(*model.Role) error
	Update(id uint, role *model.Role) error
	Delete(ids vo.RoleDeleteReq) error
	List(*vo.RoleListReq) ([]model.Role, int64, error)
}

type RoleService struct {
}

func NewRoleService() RoleServiceInterface {
	return RoleService{}
}

func (r RoleService) Create(role *model.Role) error {
	err := global.Db.Create(role).Error
	return err
}

func (r RoleService) Update(id uint, role *model.Role) error {
	err := global.Db.Model(&model.Role{}).Where("id = ?", id).Updates(role).Error
	return err
}
func (r RoleService) Delete(ids vo.RoleDeleteReq) error {
	err := global.Db.Where("id in (?)", ids.Ids).Delete(&model.Role{}).Error
	return err
}

func (r RoleService) List(req *vo.RoleListReq) ([]model.Role, int64, error) {
	list := make([]model.Role, 0)
	db := global.Db.Model(&model.Role{}).Order("created_at desc")
	if strings.TrimSpace(req.Name) != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if strings.TrimSpace(req.Keyword) != "" {
		db = db.Where("keyword like ?", "%"+req.Keyword+"%")
	}
	total := int64(0)
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	pageNum := int(req.PageNum)
	pageSize := int(req.PageSize)
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
