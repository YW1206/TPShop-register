package controller

import (
	"tpshop/dao"
	"html/template"
	"net/http"
	_ "strconv"
)

//Regist 处理用户的函注册数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	phone := r.PostFormValue("phone")
	rec_phone := r.PostFormValue("rec_phone")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist_failed.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		phonenum, _ := dao.CheckPhoneNum(phone)
		if phonenum.ID > 0 {
			t := template.Must(template.ParseFiles("views/pages/user/regist_failed.html"))
			t.Execute(w, "该手机号已存在！")
		} else {
			uniqueEmail, _ := dao.CheckEmail(email)
			if uniqueEmail.ID > 0 {
				t := template.Must(template.ParseFiles("views/pages/user/regist_failed.html"))
				t.Execute(w, "该邮箱已被注册！")		
			} else {
				//用户名可用，将用户信息保存到数据库中
				dao.SaveUser(username, password, email, phone, rec_phone)
				//用户名和密码正确
				t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
				t.Execute(w, "")
			}
		}
	}
}

//CheckUserName 通过发送Ajax验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}

func CheckPhoneNum(w http.ResponseWriter, r *http.Request) {
	phone := r.PostFormValue("phone")
	user, _ := dao.CheckPhoneNum(phone)
	if user.ID > 0 {
		w.Write([]byte("该手机号已经注册！"))
	} else {
		w.Write([]byte("<font style='color:green'>手机号可用！</font>"))
	}
}

func CheckEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	user, _ := dao.CheckEmail(email)
	if user.ID > 0 {
		w.Write([]byte("该邮箱已经注册！"))
	} else {
		w.Write([]byte("<font style='color:green'>该邮箱可用！</font>"))
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/pages/user/welcome.html"))
	t.Execute(w, "")
}
