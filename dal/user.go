package dal

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	username string
	password string
	info string
}

func SetUser(username string,key string, value string){
	var user User
	db_conn.First(&user, "username = ?", username)
	db_conn.Model(&user).Update(key, value)
}

func CreateUser(username string, password string,info string){
	db_conn.Create(&User {username: username, password: password, info: info})
}

