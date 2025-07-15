package v1

import (
	"janx-admin/app/controller"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitRoleApi(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRouter {
	userController := controller.NewRoleController()
	router := r.Group("/role")
	{
		router.POST("/create", userController.Create)
		router.PATCH("/update", userController.Update)
		router.DELETE("/delete", userController.Delete)
		router.GET("/list", userController.List)
	}
	return r
}
