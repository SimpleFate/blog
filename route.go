package main

type Router struct {
	routerMap map[string]interface{}
}

func (router *Router) Init() {
	router.routerMap["/static/"] = hStaticFile
	//router.routerMap["/"] =
}
