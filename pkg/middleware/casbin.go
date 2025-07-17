package middleware

import (
	"janx-admin/app/response"
	"janx-admin/app/service"
	"janx-admin/global"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var checkLock sync.Mutex

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		us := service.NewUserService()
		user, err := us.GetCurrentUser(c)
		if err != nil {
			response.Response(c, 401, 401, nil, "用户未登录")
			c.Abort()
			return
		}
		if user.Status != 1 {
			response.Response(c, 401, 401, nil, "用户已被禁用")
			c.Abort()
			return
		}
		var subs []string
		roles := user.Roles
		for _, role := range roles {
			if role.Status == 1 {
				subs = append(subs, role.Keyword)
			}
		}
		obj := strings.TrimPrefix(c.FullPath(), "/"+global.Conf.System.UrlPathPrefix)
		act := c.Request.Method

		isPass := check(subs, obj, act)
		if !isPass {
			response.Response(c, 401, 401, nil, "没有权限")
			c.Abort()
			return
		}
	}
}

func check(subs []string, obj string, act string) bool {
	// 同一时间只允许一个请求执行校验, 否则可能会校验失败
	checkLock.Lock()
	defer checkLock.Unlock()
	isPass := false
	for _, sub := range subs {
		pass, _ := global.CasbinEnforcer.Enforce(sub, obj, act)
		if pass {
			isPass = true
			break
		}
	}
	return isPass
}
