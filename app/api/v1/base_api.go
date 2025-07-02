package v1

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitBaseApi(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRouter {
	// userController := controller.NewUserController()
	router := r.Group("/base")
	{
		router.POST("/login", authMiddleware.LoginHandler)
		router.POST("/logout", authMiddleware.LogoutHandler)
		router.POST("/refreshToken", authMiddleware.RefreshHandler)
	}
	return r
}
