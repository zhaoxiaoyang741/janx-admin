package initializa

import (
	"fmt"
	v1 "janx-admin/app/api/v1"
	"janx-admin/global"
	"janx-admin/pkg/middleware"

	_ "janx-admin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// })
func InitRoutes(r *gin.Engine) {
	// 这里预留v1版本，之后如果更新单个接口时，切换到v2版本
	apiGroup := r.Group(global.Conf.System.UrlPathPrefix)
	apiGroup.Use(middleware.TimeMiddleware())
	// v1.InitBaseApi(apiGroup)
	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.InitUserApi(apiGroup, authMiddleware)
	v1.InitBaseApi(apiGroup, authMiddleware)
}
