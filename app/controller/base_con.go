package controller

import (
	"janx-admin/app/service"

	"github.com/gin-gonic/gin"
)

type BaseControllerInterface interface {
	Ping(c *gin.Context)
}

type BaseController struct {
	BaseService service.BaseServiceInterface
}

func NewBaseController() *BaseController {
	return &BaseController{
		BaseService: service.NewBaseService(),
	}
}

func (b *BaseController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
