package view

import (
	"fmt"
	authenticate "myweb/auth"
	"myweb/service"
	"net/http"
	"time"
)

// LoginPage 登录界面信息
type LoginPage struct {
	UserID  string
	UserPwd string
	ErrMsg  string
}

// 登录界面
func index(response http.ResponseWriter, request *http.Request) {

	renderTempalate(response, "static/html/index.html", LoginPage{})
}

// 登录处理
func login(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	userID := request.Form.Get("userid")
	password := request.Form.Get("password")

	if userID == "" || password == "" {
		renderTempalate(response, "static/html/index.html", LoginPage{
			UserID: userID,
			ErrMsg: "请输入用户ID和密码",
		})
		return
	}

	userInfo, err := service.GetUserInfo(userID)

	if err != nil {
		renderTempalate(response, "static/html/index.html", LoginPage{
			UserID: userID,
			ErrMsg: "数据库异常",
		})
		return
	}

	if userInfo == nil || userInfo.UserPwd != password {
		renderTempalate(response, "static/html/index.html", LoginPage{
			UserID: userID,
			ErrMsg: "用户ID或密码错误",
		})
		return
	}

	// 发行Token并写入cookie
	returnCode, err := authenticate.DistributeToken(response, userID)
	response.WriteHeader(returnCode)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("User: %v(%v)\n", userInfo.UserName, userID)
	fmt.Printf("Time: %v\n", time.Now())
	fmt.Printf("IP: %v\n", request.RemoteAddr)

	renderTempalate(response, "static/html/main.html", nil)
}
