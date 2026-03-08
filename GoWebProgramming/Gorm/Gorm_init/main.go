package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局数据库连接（后续所有操作都用这个db）
var db *gorm.DB

func init() {
	// 1. 拼接DSN（替换成你的MySQL配置）
	dsn := "root:0190360243Zsong/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	// 2. 连接数据库
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接MySQL失败：" + err.Error())
	}
}

func main() {
	// 后续写CRUD代码
}
