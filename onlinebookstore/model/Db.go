// 数据库的入口文件，连接、配置数据库
package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"onlinebookstore/utils"
	"time"
)

var(
	db *gorm.DB
	err error
)

func InitDb(){
	// 连接数据库
 	dsn := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
 	db , err = gorm.Open(dsn, &gorm.Config{})
	if err != nil{
		fmt.Println("连接数据库失败，请检查数据库连接的参数：",err)
	}

	//GORM 提供了 DB 方法，可用于从当前 *gorm.DB 返回一个通用的数据库接口 *sql.DB
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// 迁移 AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的
	// AutoMigrate 会创建表，缺少的外键，约束，列和索引，并且会更改现有列的类型
	//（如果其大小、精度、是否为空可更改）。 但不会删除未使用的列，以保护您的数据。
	//db.AutoMigrate(&User{},&Category{},&Book{},Personal{})
	db.AutoMigrate(&User{},&Category{},&Book{}, &Cart{})


	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)


	//sqlDB.Close()
}