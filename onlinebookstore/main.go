package main

import (
	"onlinebookstore/model"
	"onlinebookstore/routes"
)

//go mod tidy 用于go mod 依赖加载

func main(){
	// 引用数据库
	model.InitDb()
	// 引用路由
	routes.InitRouter()
}