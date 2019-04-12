package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	_"github.com/mattn/go-sqlite3"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "welcome to book-recommend!")
}

func UserRegister(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	info := r.Form.Get("info")

}

func main(){
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/user.register", UserRegister)
	err := http.ListenAndServe(":8890", nil)

	if err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}