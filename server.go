package main

import (
	"blog/consts/serverconst"
	"blog/db"
	"fmt"
	"log"
	"net/http"
)

func startServer() {
	db.Init()
	defer db.Destroy()
	addr := fmt.Sprintf(":%d", serverconst.Port)
	err := http.ListenAndServe(addr, nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
