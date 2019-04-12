package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


func get_conn()(db *gorm.DB, err error){
	db, err = gorm.Open("sqlite3", "./book-commend.db")
	if err != nil{
		log.Fatal(err)
		return db, err
	}
	return db, nil
}

var db_conn *gorm.DB

func init(){
	db_conn, _ = get_conn()
	log.Println("check database  init")
	if !db_conn.HasTable(&User{}){
		db_conn.CreateTable(&User{})
		log.Println("user table create success")
	}
}