package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


func get_conn()(db *gorm.DB, err error){
	db, err = gorm.Open("sqlite3", "db/book-commend.db")
	if err != nil{
		log.Fatal(err)
		return db, err
	}
	return db, nil
}

var db_conn *gorm.DB

func __int__(){
	db_conn, _ = get_conn()
	if !db_conn.HasTable(&User{}){
		db_conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}
}