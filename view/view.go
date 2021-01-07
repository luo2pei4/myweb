package view

import (
	"net/http"
	"text/template"
)

// RegistRoutine 注册路由
func RegistRoutine() {

	// 声明静态资源路径
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 配置其他路径的方法
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/choice", choice)
	http.HandleFunc("/searchExRate", searchExRate)
	http.HandleFunc("/adsbCoverRange", adsbCoverRange)
}

// RenderTempalate 渲染网页
func renderTempalate(response http.ResponseWriter, filePath string, data interface{}) {
	t := template.Must(template.ParseFiles(filePath))
	t.Execute(response, data)
}

// 相应菜单栏请求, 跳转到不同界面.
func choice(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	fid := request.Form.Get("fid")

	switch fid {
	case "1":
		renderTempalate(response, "static/html/exchange.html", nil)
	case "2":
		renderTempalate(response, "static/html/adsbCoverRange.html", nil)
	default:
		response.Write([]byte("maintainance"))
	}
}
