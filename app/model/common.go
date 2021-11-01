package model

import (
	"gorm.io/gorm"
	"scut2022-bishe/util"
)

var DB *gorm.DB

func LoadModelDB() {
	DB = util.GetDB()
}
