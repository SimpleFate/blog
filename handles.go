package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var (
	hStaticFile = http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP
	hTest       = sayhelloName
	rIndex      = index
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}

func index(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("view/learn02.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = res.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
