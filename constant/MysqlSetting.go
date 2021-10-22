package constant

import (
	"gopkg.in/ini.v1"
	"scut2022-bishe/app/middleware"
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
		middleware.Logger().Errorf("[mysql]配置文件加载错误, %s", err)
	}

	err = cfg.Section("mysql").MapTo(MysqlSetting)
	if err != nil {
		middleware.Logger().Errorf("[mysql]配置文件映射错误, %s", err)
	}
}
