package global

import (
	"janx-admin/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	SysViper *viper.Viper
	Conf     config.Config
	Db       *gorm.DB
	Logger   *logrus.Logger
)
