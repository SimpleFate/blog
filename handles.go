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
)
var (
	handleAddComment     = apiAddComment
	handleSupportComment = apiSupportComment
	handleOpposeComment  = apiOpposeComment
	handleReplyComment   = apiReplyComment
	handleListComments   = apiListComments
	handleGetReplys      = apiGetReplys
)

//留言板

func apiAddComment(w http.ResponseWriter, r *http.Request) {
	param := struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}{}
	err := getJsonFromBody(r, &param)
	if err != nil {
		fmt.Println(err)
		return
	}

	privacy := services.GetPrivacy(r)
	services.AddComment(privacy, param.Name, param.Content)
}
func apiSupportComment(w http.ResponseWriter, r *http.Request) {
	param := struct {
		Id string `json:"id"`
	}{}
	err := getJsonFromBody(r, &param)
	if err != nil {
		fmt.Println(err)
		return
	}
	services.SupportComment(param.Id)
}
func apiOpposeComment(w http.ResponseWriter, r *http.Request) {
	param := struct {
		Id string `json:"id"`
	}{}
	err := getJsonFromBody(r, &param)
	if err != nil {
		fmt.Println(err)
		return
	}
	services.OpposeComment(param.Id)
}
func apiReplyComment(w http.ResponseWriter, r *http.Request) {
	param := struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Remark string `json:"remark"`
	}{}
	err := getJsonFromBody(r, &param)
	if err != nil {
		fmt.Println(err)
		return
	}
	privacy := services.GetPrivacy(r)
	services.ReplyComment(param.Id, param.Name, param.Remark, privacy)
}

func apiListComments(w http.ResponseWriter, r *http.Request) {
	res := services.ListTopComments()
	jsonRes, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonRes)
}

func apiGetReplys(w http.ResponseWriter, r *http.Request) {
	param := struct {
		Id string `json:"id"`
	}{}
	err := getJsonFromBody(r, &param)
	if err != nil {
		return
	}
	res := services.GetReplys(param.Id)
	jsonRes, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonRes)
}

//获取body中的json
//result 自定义结构体
func getJsonFromBody(r *http.Request, result interface{}) error {
	body := r.Body
	defer body.Close()

	buffers, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffers, result)
	if err != nil {
		return err
	}
	return nil
}

//跳转。废弃，使用Vue跳转
func forward(w http.ResponseWriter, r *http.Request, f string) {
	hanle := getForwardHandle(f)
	hanle(w, r)
}

//获取页面跳转HandleFunc。废弃，使用Vue跳转
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
