package global

import (
	"janx-admin/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	SysViper *viper.Viper
	Conf     config.Config
	Db       *gorm.DB
	Logger   *logrus.Logger

	// 全局Validate数据校验实列
	Validate *validator.Validate

	// 全局翻译器
	Trans ut.Translator
)
