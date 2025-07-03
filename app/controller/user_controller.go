package controller

import (
	"janx-admin/app/model"
	"janx-admin/app/response"
	"janx-admin/app/service"
	"janx-admin/app/vo"
	"janx-admin/global"
	encUtils "janx-admin/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserServiceInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController() UserServiceInterface {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// 创建用户
// @Summary 创建用户
// @Description  创建新用户
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body vo.UserCreateReq true "用户创建请求参数"
// @Success 200 {object} response.ResponseData{message=string} "创建用户成功"
// @Router /api/v1/user/create [post]
func (uc *UserController) Create(c *gin.Context) {
	var req vo.UserCreateReq
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	//参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	user := &model.User{}
	password := req.Password
	if password == "" {
		password = "123456"
	}

	encPassword, err := encUtils.EncryptPassword(password)
	if err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}
	user.Username = req.Username
	user.Nickname = req.Nickname
	user.Password = encPassword
	user.Email = req.Email
	user.Phone = req.Phone
	user.Avatar = req.Avatar
	user.CreatedAt = time.Now()
	user.Status = 1
	err = uc.userService.Create(user)
	if err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}
	response.Success(c, nil, "创建用户成功")
}

// 更新用户信息
// @Summary 更新用户信息
// @Description  更新用户信息
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body vo.UserUpdataReq true "用户更新请求参数"
// @Success 200 {object} response.ResponseData{message=string} "更新用户成功"
// @Router /api/v1/user/update [patch]
func (uc *UserController) Update(c *gin.Context) {
	var req vo.UserUpdataReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	//参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	user := &model.User{}
	user.Avatar = req.Avatar
	user.Nickname = req.Nickname
	user.Email = req.Email
	user.Phone = req.Phone
	err := uc.userService.Update(req.ID, user)
	if err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}
	response.Success(c, nil, "更新用户成功")

}

// 删除用户信息
// @Summary 删除用户信息
// @Description  根据id删除用户信息
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body vo.UserDeleteReq true "用户删除请求参数"
// @Success 200 {object} response.ResponseData{message=string} "删除用户成功"
// @Router /api/v1/user/delete [delete]
func (uc *UserController) Delete(c *gin.Context) {
	var req vo.UserDeleteReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	//参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	err := uc.userService.Delete(req)
	if err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}
	response.Success(c, nil, "删除用户成功")
}

// 获取用户信息列表
// @Summary 获取用户信息列表
// @Description  获取用户信息列表(username 和 nickname 可模糊查询)
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body vo.UserListReq false "用户列表获取请求参数"
// @Success 200 {object} response.ResponseData{data=response.ResListData{list=[]model.User}} "成功返回用户列表"
// @Router /api/v1/user/list [get]
func (uc *UserController) List(c *gin.Context) {
	var req vo.UserListReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	//参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		response.Fail(c, nil, errStr)
		return
	}
	userList, total, err := uc.userService.List(&req)
	if err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"list":  userList,
		"total": total,
	}, "获取用户列表成功")
}
