package models

import (
	"shuaibingBlog/utils"
	"time"
)

type User struct {
	Id           int64  //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Nickname     string //`name` varchar(32) COMMENT '用户名',
	Email        string //`email` varchar(32) DEFAULT NULL COMMENT '用户邮箱',
	Password     string //`password` varchar(32) COMMENT '用户密码',
	Gender       int64  //`mobile` varchar(20) DEFAULT NULL COMMENT '用户手机号',
	Created_time time.Time
}

// 数据表名称
func (this *User) TableName() string {
	return "user"
}

func init() {
	// 注册定义的model
	utils.SqlDB.AutoMigrate(&User{})
}

type UserModel struct {
	//业务模型
}

func (this *UserModel) AddOne(u User) (err error) {
	utils.SqlDB.NewRecord(&u)
	return utils.SqlDB.Error
}

func (this *UserModel) FindByName(user *User) (err error) {
	utils.SqlDB.First(&user, "nickname = ?", user.Nickname)
	err = utils.SqlDB.Error
	return
}
