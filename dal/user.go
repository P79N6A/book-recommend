package dal

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Info string `gorm:"column:info"`
}

func SetUser(username string,key string, value string){
	var user User
	db_conn.First(&user, "username = ?", username)
	db_conn.Model(&user).Update(key, value)
}

func CreateUser(username string, password string,info string){
	db_conn.Create(&User {Username: username, Password: password, Info: info})
}

func CheckUserExist(username string) bool {
	var user User
	isNotFound := db_conn.First(&user, "username = ?", username).RecordNotFound()
	if isNotFound {
		return false
	}
	return true
}

func CheckUserPassword(username string, password string)(bool, string){
	var user User
	isNotFound := db_conn.First(&user, "username = ?", username).RecordNotFound()
	if isNotFound {
		return false, ""
	}
	return true, user.Info

}
