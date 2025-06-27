package v1

import (
	"janx-admin/app/controller"

	"github.com/gin-gonic/gin"
)

func InitUserApi(r *gin.RouterGroup) gin.IRouter {
	userController := controller.NewUserController()
	router := r.Group("/user")
	{
		router.GET("/create", userController.Create)
	}
	return r
}
