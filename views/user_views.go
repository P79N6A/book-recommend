package views

import (
	"book-recommend/dal"
	"code.byted.org/gopkg/pkg/log"
	"encoding/json"
	"fmt"
	"net/http"
)

type responseJson struct {
	Success bool `json:"success"`
	Msg string `json:"msg"`
}

func Register(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	log.Info(r.Form)
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	info := r.Form.Get("info")

	if dal.CheckUserExist(username){
		res := responseJson{Success:false, Msg:msgText["register_failed"]}
		resStr, _ := json.Marshal(res)

		log.Info(string(resStr))
		fmt.Fprint(w, string(resStr))
		return
	}
	dal.CreateUser(username, password, info)
	res := responseJson{Success:true, Msg:msgText["register_success"]}
	resStr, _ := json.Marshal(res)

	log.Info(string(resStr))
	fmt.Fprint(w, string(resStr))
}

func Login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	log.Info(r.Form)
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	success, userInfo := dal.CheckUserPassword(username, password)
	var res map[string]interface{}
	res = make(map[string]interface{})
	if not success{
		res["success"] = false
		res["msg"] = msgText["login_failed"]
		res["data"] = ""
	}else {
		res["success"] = true
		res["msg"] = msgText["login_success"]
		res["data"] = userInfo
	}
	resJsonText, _ := json.Marshal(res)

	log.Info(string(resJsonText))
	fmt.Fprint(w, string(resJsonText))
}

func ModifyUserInfoOrInfo(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	log.Info(r.Form)
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	info := r.Form.Get("info")

	if dal.ModifyUser(username){
		res := responseJson{Success:false, Msg:msgText["modify_failed"]}
		resStr, _ := json.Marshal(res)

		log.Info(string(resStr))
		fmt.Fprint(w, string(resStr))
		return
	}
	dal.CreateUser(username, password, info)
	res := responseJson{Success:true, Msg:msgText["modify_success"]}
	resStr, _ := json.Marshal(res)

	log.Info(string(resStr))
	fmt.Fprint(w, string(resStr))
}













