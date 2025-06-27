package middleware

import (
	"fmt"
	"janx-admin/app/model"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := model.SystemRouteLog{}
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		// 计算耗时
		elapsedTime := endTime.Sub(startTime)
		req.Url = c.Request.URL.Path
		req.Method = c.Request.Method
		req.Duration = int64(elapsedTime)
		req.Ip = c.ClientIP()
		req.UserAgent = c.Request.UserAgent()
		fmt.Printf("请求路径：%v\n", req)
	}
}
