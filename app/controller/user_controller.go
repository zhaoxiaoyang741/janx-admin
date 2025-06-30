package controller

import (
	"janx-admin/app/controller/utils"
	"janx-admin/app/model"
	"janx-admin/app/service"
	"janx-admin/app/vo"
	encUtils "janx-admin/pkg/utils"
	"time"

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
	var r vo.UserCreateReq
	err := utils.CheckReq(c, &r)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	user := &model.User{}
	password := r.Password
	if password == "" {
		password = "123456"
	}

	encPassword, err := encUtils.EncryptPassword(password)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	user.Username = r.Username
	user.Nickname = r.Nickname
	user.Password = encPassword
	user.Email = r.Email
	user.Phone = r.Phone
	user.Avatar = r.Avatar
	user.CreatedAt = time.Now()
	user.Status = 1
	err = uc.userService.Create(*user)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	utils.Success(c, nil, "创建用户成功")
}

func (uc *UserController) Update(c *gin.Context) {
	var r vo.UserUpdataReq
	err := utils.CheckReq(c, &r)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	user := &model.User{}
	user.Avatar = r.Avatar
	user.Nickname = r.Nickname
	user.Email = r.Email
	user.Phone = r.Phone
	err = uc.userService.Update(r.ID, *user)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	utils.Success(c, nil, "更新用户成功")

}

func (uc *UserController) Delete(c *gin.Context) {
	var r vo.UserDeleteReq
	err := utils.CheckReq(c, &r)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	err = uc.userService.Delete(r)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	utils.Success(c, nil, "删除用户成功")
}

func (uc *UserController) List(c *gin.Context) {
	var r vo.UserListReq
	err := utils.CheckReq(c, &r)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	userList, total, err := uc.userService.List(&r)
	if err != nil {
		utils.FailWithoutData(c, err.Error())
		return
	}
	utils.Success(c, gin.H{
		"list":  userList,
		"total": total,
	}, "获取用户列表成功")
}
