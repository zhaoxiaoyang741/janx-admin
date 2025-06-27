package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	System System `json:"system" yaml:"system"`
	Db     Db     `json:"db" yaml:"db"`
	Log    Log    `json:"log" yaml:"log"`
}

func InitConfig() Config {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("config init fail : %s", err.Error()))
	}
	fmt.Printf("Read config file: %s is successful\n", "./resources/config.yaml")

	// 读取配置到 global.Conf
	conf := Config{}
	err = v.Unmarshal(&conf)
	if err != nil {
		panic("unmarshal config file is fail")
	}

	fmt.Println("Config file data:", v.AllSettings())
	// 监听变化
	// v.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// 	err = v.Unmarshal(&conf)
	// 	if err != nil {
	// 		fmt.Println("Failed to unmarshal updated config:", err)
	// 	}
	// })

	// v.WatchConfig()
	return conf
}
