package main

import "net/http"

var (
	routerMap = make(map[string]http.HandlerFunc)
)

func init() {
	routerMap["/"] = pageIndex
	routerMap["/static/"] = handleStatics

	routerMap["/api/comment/add"] = handleAddComment
	bind()
}
func bind() {
	for k, v := range routerMap {
		http.HandleFunc(k, v)
	}
}
