package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	initializa "janx-admin/Initializa"
	"janx-admin/config"
	"janx-admin/global"
)

// @title jiax-admin
// @version 1.0
// @description jiax-admin is a private management background
// @termsOfService http://swagger.io/terms/

// @contact.name cwy
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host zhaoxiaoyang.com
// @BasePath /api
func main() {
	// 初始化配置
	global.Conf = config.InitConfig()

	// 初始化日志
	initializa.InitLogger()

	// 初始化mysql
	initializa.InitPostgres()

	// 初始化http service
	initializa.InitService()

	// 初始化
	initializa.InitValidate()

	// 设置信号处理，使主程序保持运行
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("接收到终止信号，程序退出.")
}
