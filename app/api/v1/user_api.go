package v1

import (
	"janx-admin/app/controller"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitUserApi(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRouter {
	userController := controller.NewUserController()
	router := r.Group("/user")
	{
		router.POST("/create", userController.Create)
		router.PATCH("/update", userController.Update)
		router.DELETE("/delete", userController.Delete)
		router.GET("/list", userController.List)
	}
	return r
}
