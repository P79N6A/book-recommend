package views

import "log"

var msgText  map[string]string


func init(){
	msgText = make(map[string]string)
	msgText["register_success"] = "注册成功"
	msgText["register_failed"] = "注册失败，用户名已存在或信息填写不规范"
	msgText["login_success"] = "登陆成功"
	msgText["login_failed"] = "登陆失败，请检查用户名或密码"

	log.Println("text const init success")
}

