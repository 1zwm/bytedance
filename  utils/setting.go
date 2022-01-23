package Utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	Db string
	DbPort string
	DbHost string
	DbUser string
	DbPassWord string
	DbName string
)

func init() {
	file, err := ini.Load("../config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查读取路径 ：", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File)  {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8090")
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("20011126")
	DbName = file.Section("database").Key("DbName").MustString("bytedance")
}