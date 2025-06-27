package v1

import (
	"janx-admin/app/controller"

	"github.com/gin-gonic/gin"
)

func InitBaseApi(r *gin.RouterGroup) gin.IRouter {
	baseController := controller.NewBaseController()

	router := r.Group("/base")
	{
		router.GET("/ping", baseController.Ping)
	}
	return r
}
