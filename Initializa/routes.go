package initializa

import (
	v1 "janx-admin/app/api/v1"
	"janx-admin/global"
	"janx-admin/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// })
func InitRoutes(r *gin.Engine) {
	// 这里预留v1版本，之后如果更新单个接口时，切换到v2版本
	apiGroup := r.Group(global.Conf.System.UrlPathPrefix)
	apiGroup.Use(middleware.TimeMiddleware())
	// v1.InitBaseApi(apiGroup)
	v1.InitUserApi(apiGroup)
}
