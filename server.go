package main

import (
	"blog/consts/serverconst"
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	addr := fmt.Sprintf(":%d", serverconst.Port)
	err := http.ListenAndServe(addr, nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
