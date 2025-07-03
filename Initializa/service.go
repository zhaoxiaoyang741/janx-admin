package initializa

import (
	"fmt"
	"janx-admin/global"

	"github.com/gin-gonic/gin"
)

func InitService() {
	//是否开启debug
	if !global.Conf.System.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 设置可信代理，避免安全风险
	r.SetTrustedProxies([]string{"127.0.0.1"})

	InitRoutes(r)

	address := fmt.Sprintf("127.0.0.1:%d", global.Conf.System.Port)

	go func() {
		if err := r.Run(address); err != nil {
			global.Logger.Error("HTTP服务启动失败: " + err.Error())
		}
	}()

	// global.Logger.Info(fmt.Sprintf("HTTP服务已启动，运行地址：http://%s\n", address))
	fmt.Printf("HTTP服务已启动，运行地址：http://%s\n", address)
	global.Logger.Info(fmt.Sprintf("swagger文档地址：http://%s/swagger/index.html", address))
}
