package main

import (
	"fmt"
	"tpshop/controller"
	"net/http"
)

func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//去首页
	http.HandleFunc("/main", controller.GetUserInfo)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有用户
	http.HandleFunc("/getUserInfo", controller.GetUserInfo)

	fmt.Println("系统已经启动，访问系统的地址为","http://127.0.0.1:8080/main")

	http.ListenAndServe(":8080", nil)
}
