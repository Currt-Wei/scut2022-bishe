package main

import (
	"scut2022-bishe/constant"
	"scut2022-bishe/util"
)

func init() {
	constant.InitMysqlSetting()
	util.InitMysql()
}
