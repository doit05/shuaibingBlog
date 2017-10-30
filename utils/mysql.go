package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"shuaibingBlog/helper"
)

var SqlDB *gorm.DB

func init() {
	conf := helper.Config
	db, err := gorm.Open("mysql", conf.Mysqlcon)
	if err != nil {
		panic("设置DB数据库配置失败, err: " + err.Error())
	}
	SqlDB = db
}
