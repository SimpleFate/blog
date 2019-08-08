package main

import (
	"blog/services"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var (
	viewFilePrefix = "view/"
)

var (
	handleStatics = http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP
	pageIndex     = getForwardHandle("learn02.html")
)
var (
	handleAddComment = apiAddComment
)

//留言板

func apiAddComment(w http.ResponseWriter, r *http.Request) {
	param := struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}{}
	getJsonFromBody(r, &param)

	privacy := services.GetPrivacy(r)
	services.AddComment(privacy, param.Name, param.Content)
}
func apiSupportComment(w http.ResponseWriter, r *http.Request) {
	param := struct {
		Id string `json:"id"`
	}{}
	getJsonFromBody(r, &param)

}
func apiOpposeComment(w http.ResponseWriter, r *http.Request) {

}

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//parm := struct {
	//	Key string `json:"key"`
	//}{}

	//getJsonFromBody(r, &parm)

	//关闭xss保护，可以执行脚本
	//w.Header().Set("X-XSS-Protection", "0")

}

func PrintHeader(header http.Header) {
	for k, vs := range header {
		fmt.Printf("%s: ", k)
		for _, v := range vs {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}
}

//获取body中的json
//result 自定义结构体
func getJsonFromBody(r *http.Request, result interface{}) {
	body := r.Body
	defer body.Close()
	//
	//buffer := make([]byte, 1024)
	//buffers := make([]byte, 0, 1024)
	//n, _ := body.Read(buffer)
	//for n > 0 {
	//	//str := (*string)(unsafe.Pointer(&buffer))
	//	buffers = append(buffers, buffer[:n]...)
	//	n, _ = body.Read(buffer)
	//}

	buffers, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(buffers, result)
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
