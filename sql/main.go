package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/moniang/chat/config"
)

var DB *gorm.DB

// 初始化MYSQL连接
func InitDb() {
	var err error
	DB, err = gorm.Open("mysql", config.DbUser+":"+config.DbPass+"@("+config.DbAddr+")/"+config.DbName+"?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(config.DbDebug)
	if err != nil {
		fmt.Println(err)
	}
}
