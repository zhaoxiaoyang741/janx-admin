package controller

import (
	"fmt"
	"janx-admin/app/model"
	"janx-admin/app/service"
	"janx-admin/global"

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	Create(*gin.Context)
}

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) Create(c *gin.Context) {
	user := &model.User{
		Username: "admin",
		Password: "admin",
		Phone:    "13389057029",
		Email:    "admin@admin.com",
	}
	err := uc.userService.Create(*user)
	global.Logger.Info(fmt.Sprintf("create user err:%v", err))
}
