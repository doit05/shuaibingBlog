package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var SqlDB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:root123@tcp(10.118.26.144:3306)/goweb_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("设置Test_db数据库配置失败, err: " + err.Error())
	}
	SqlDB = db
}
