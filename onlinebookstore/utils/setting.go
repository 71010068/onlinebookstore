package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// 将配置文件中数据定义成全局变量
var(
	AppMode 	string
	HttpPort 	string
	JwtKey		string

	Db 			string
	DbHost 		string
	DbPort 		string
	DbUser		string
	DbPassWord 	string
	DbName 		string

	AccessKey	string
	SecretKey	string
	Bucket		string
	QiniuServer	string
)
// init 初始化ini中的数据
func init(){
	file , err := ini.Load("config/config.ini")
	if err != nil{
		fmt.Println("配置文件读取错误，请检查文件路径。" , err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debuge")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("HttpPort").MustString("520lxg1314")
}
func LoadData(file *ini.File){
	Db 		 = file.Section("database").Key("Db").MustString("mysql")
	DbHost 	 = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort 	 = file.Section("database").Key("DbPort").MustString("3306")
	DbUser	 = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord  = file.Section("database").Key("DbPassWord").MustString("0415")
	DbName 	 = file.Section("database").Key("DbName").MustString("onlinebookstore")
}
func LoadQiniu(file *ini.File){
	AccessKey	=	file.Section("qiniu").Key("AccessKey").String()
	SecretKey	=	file.Section("qiniu").Key("SecretKey").String()
	Bucket	=	file.Section("qiniu").Key("Bucket").String()
	QiniuServer	=	file.Section("qiniu").Key("QiniuServer").String()
}