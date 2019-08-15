package main

import "net/http"

var (
	routerMap = make(map[string]http.HandlerFunc)
)

func init() {
	//routerMap["/"] = pageIndex
	routerMap["/static/"] = handleStatics

	//post
	routerMap["/api/comment/add"] = handleAddComment
	routerMap["/api/comment/support"] = handleSupportComment
	routerMap["/api/comment/oppose"] = handleOpposeComment
	routerMap["/api/comment/reply/add"] = handleReplyComment
	routerMap["/api/comment/reply/get"] = handleGetReplys

	//get
	routerMap["/api/comment/list"] = handleListComments

	bind()
}
func bind() {
	for k, v := range routerMap {
		http.HandleFunc(k, v)
	}
}
