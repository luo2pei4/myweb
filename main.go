package main

import (
	"log"
	cfgloader "myweb/config"
	"myweb/dao"
	"myweb/view"
	"net/http"
)

func init() {

	// 加载配置文件
	cfgloader.LoadConfig("config.toml")
}

func main() {

	// 加载数据库连接
	dao.LoadConnections()

	// 注册Request路由
	view.RegistRoutine()

	//设置监听的端口
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
