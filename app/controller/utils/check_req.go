package utils

import (
	"janx-admin/global"

	"github.com/gin-gonic/gin"
)

func CheckReq(c *gin.Context, req any) error {

	// 绑定请求参数
	if err := c.ShouldBind(req); err != nil {
		return err
	}

	// 验证请求参数
	if err := global.Validate.Struct(req); err != nil {
		// err:= errors.New(err.Error())
		// errstr := err.(validator.ValidationErrors)[0].Translate(global.Trans)
		// FailWithoutData(c, fmt.Sprintf("参数校验失败失败:%s", errstr))
		// FailWithoutData(c, "参数校验失败")
		return err
	}
	return nil

}
