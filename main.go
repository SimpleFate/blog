package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP)

	http.HandleFunc("/", index)            // 设置访问的路由
	err := http.ListenAndServe(":80", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
