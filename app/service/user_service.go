package service

import (
	"janx-admin/app/model"
	"janx-admin/global"
)

type UserServiceInterface interface {
	Create(model.User) error
}

type UserService struct {
}

func NewUserService() UserServiceInterface {
	return &UserService{}
}

func (u *UserService) Create(user model.User) error {
	err := global.Db.Create(&user).Error
	return err
}
