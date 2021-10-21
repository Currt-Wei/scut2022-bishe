package constant

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Mysql struct {
	Ip       string `ini:"ip"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

var MysqlSetting = &Mysql{}

func InitMysqlSetting() {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("读取配置错误")
	}

	err = cfg.Section("mysql").MapTo(MysqlSetting)
	if err != nil {
		fmt.Println("ini文件映射错误")
	}
}
