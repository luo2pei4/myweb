package main

import (
	"log"
	"myweb/view"
	"net/http"
)

func main() {

	view.RegistRoutine()

	//设置监听的端口
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
