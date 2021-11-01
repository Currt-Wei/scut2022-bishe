package main

import (
	"scut2022-bishe/app/model"
	"scut2022-bishe/constant"
	"scut2022-bishe/util"
	"scut2022-bishe/util/casbin"
)

func init() {
	constant.InitMysqlSetting()
	util.InitMysql()
	model.LoadModelDB() // 加载model中使用的db

	// 加载casbin策略
	casbin.InitCasbinPolicyData()
}
