package controller

import (
	"janx-admin/app/model"
	"janx-admin/app/response"
	"janx-admin/app/service"
	"janx-admin/app/vo"
	"janx-admin/global"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoleControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type RoleController struct {
	roleService service.RoleServiceInterface
}

func NewRoleController() RoleControllerInterface {
	return &RoleController{
		roleService: service.NewRoleService(),
	}
}

// @Summary 创建角色
// @Description  创建角色
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param body body vo.RoleCreateReq true "角色创建请求参数"
// @Success 200 {object} response.ResponseData{message=string} "创建角色成功"
// @Router /api/v1/role/create [post]
func (r *RoleController) Create(c *gin.Context) {
	var req vo.RoleCreateReq
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

	role := &model.Role{
		Name:    req.Name,
		Keyword: req.Keyword,
		Desc:    req.Desc,
		Status:  req.Status,
		Sort:    req.Sort,
	}
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	if err := r.roleService.Create(role); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	response.Success(c, nil, "创建成功")
}

// @Summary 更新角色信息
// @Description  更新角色信息
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body vo.RoleUpdateReq true "角色更新请求参数"
// @Success 200 {object} response.ResponseData{message=string} "角色用户成功"
// @Router /api/v1/role/update [patch]
func (r *RoleController) Update(c *gin.Context) {
	var req vo.RoleUpdateReq
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

	role := &model.Role{
		Name:    req.Name,
		Keyword: req.Keyword,
		Desc:    req.Desc,
		Status:  req.Status,
		Sort:    req.Sort,
	}
	// role.ID = req.ID
	role.UpdatedAt = time.Now()

	if err := r.roleService.Update(req.ID, role); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}
}

// 删除角色信息
// @Summary 删除角色信息
// @Description  根据id删除角色信息
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body vo.RoleDeleteReq true "角色删除请求参数"
// @Success 200 {object} response.ResponseData{message=string} "删除角色成功"
// @Router /api/v1/role/delete [delete]
func (r *RoleController) Delete(c *gin.Context) {
	var req vo.RoleDeleteReq
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

	if err := r.roleService.Delete(req); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	response.Success(c, nil, "删除成功")

}

// 获取角色信息列表
// @Summary 获取角色信息列表
// @Description  获取角色信息列表(name 和 keyword 可模糊查询)
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body vo.RoleListReq false "角色列表获取请求参数"
// @Success 200 {object} response.ResponseData{data=response.ResListData{list=[]model.Role}} "成功返回角色列表"
// @Router /api/v1/user/list [get]
func (r *RoleController) List(c *gin.Context) {
	var req vo.RoleListReq
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	//参数校验
	if err := global.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		response.Fail(c, nil, errStr)
	}

	list, total, err := r.roleService.List(&req)
	if err != nil {
		response.FailWithoutData(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": list, "total": total}, "查询成功")
}
