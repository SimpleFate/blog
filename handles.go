package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var (
	viewFilePrefix = "view/"
)

var (
	handleStatics = http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP
	handleTest    = test

	pageIndex = getForwardHandle("learn02.html")
)

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//parm := struct {
	//	Key string `json:"key"`
	//}{}

	value := r.Form.Get("username")
	//getJsonFromBody(r, &parm)

	//关闭xss保护，可以执行脚本
	//w.Header().Set("X-XSS-Protection", "0")

	fmt.Fprintln(w, value)
}

//获取body中的json
//result 自定义结构体
func getJsonFromBody(r *http.Request, result interface{}) {
	body := r.Body
	defer body.Close()

	buffer := make([]byte, 1024)
	buffers := make([]byte, 0, 1024)
	n, _ := body.Read(buffer)
	for n > 0 {
		//str := (*string)(unsafe.Pointer(&buffer))
		buffers = append(buffers, buffer[:n]...)
		n, _ = body.Read(buffer)
	}
	err := json.Unmarshal(buffers, result)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//跳转
func forward(w http.ResponseWriter, r *http.Request, f string) {
	hanle := getForwardHandle(f)
	hanle(w, r)
}

//获取页面跳转HandleFunc
func getForwardHandle(templateFileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templatePath := fmt.Sprintf("%s%s", viewFilePrefix, templateFileName)
		res, err := template.ParseFiles(templatePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = res.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}
