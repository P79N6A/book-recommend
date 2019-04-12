package main

import (
	"book-recommend/views"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strings"
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


func main(){
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/user.register", views.Register)
	err := http.ListenAndServe(":8890", nil)

	if err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}