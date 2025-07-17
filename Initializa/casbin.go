package initializa

import (
	"fmt"
	"janx-admin/global"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func InitCasbinEnforcer() {
	e, err := mysqlCasbin()
	if err != nil {
		global.Logger.Panicf("初始化Casbin失败：%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}

	global.CasbinEnforcer = e
	global.Logger.Info("初始化Casbin完成!")
}

func mysqlCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDBWithCustomTable(global.Db, nil, "casbin_rule")
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(global.Conf.Casbin.ModelPath, a)
	if err != nil {
		return nil, err
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
