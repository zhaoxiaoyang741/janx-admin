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

func main() {
	// 初始化配置
	global.Conf = config.InitConfig()

	// 初始化日志
	initializa.InitLogger()

	// 初始化mysql
	initializa.InitMySQL()

	// 初始化http service
	initializa.InitService()

	// 设置信号处理，使主程序保持运行
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("接收到终止信号，程序退出.")
}
