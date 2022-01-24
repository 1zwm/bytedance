package model

import (
	"bytedance/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

//用户名：密码@tcp(Host:Port)/数据库名)

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName,
		)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}
}
