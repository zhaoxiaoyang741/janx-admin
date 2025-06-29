package service

import (
	"errors"
	"janx-admin/app/model"
	"janx-admin/app/req"
	"janx-admin/global"
	"strings"
)

type UserServiceInterface interface {
	Create(model.User) error
	Update(id uint, user model.User) error
	Delete(ids req.UserDeleteReq) error
	List(r *req.UserListReq) (list []*model.User, total int64, err error)
}

type UserService struct {
}

func NewUserService() UserServiceInterface {
	return &UserService{}
}

func (u *UserService) Create(user model.User) error {
	users := []model.User{}
	global.Db.Where("username = ? or phone = ?", user.Username, user.Phone).Find(&users)
	if len(users) > 0 {
		return errors.New("用户已存在(用户名或手机号已存在)，请勿重新添加")
	}
	err := global.Db.Create(&user).Error
	return err
}

func (u *UserService) Update(id uint, user model.User) error {
	err := global.Db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error
	return err
}

func (u *UserService) Delete(ids req.UserDeleteReq) error {
	err := global.Db.Where("id in (?)", ids.Ids).Delete(&model.User{}).Error
	return err
}

func (u *UserService) List(r *req.UserListReq) (list []*model.User, total int64, err error) {
	db := global.Db.Model(&model.User{}).Order("created_at desc")
	if strings.TrimSpace(r.Username) != "" {
		db.Where("username like ?", "%"+r.Username+"%").Find(&list)
	}
	if strings.TrimSpace(r.Nickname) != "" {
		db.Where("nickname like ?", "%"+r.Nickname+"%").Find(&list)
	}
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := int(r.PageNum)
	pageSize := int(r.PageSize)
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
