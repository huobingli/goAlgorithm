package main

import (
	"net/http"
)

// 需要建立一个文件服务器
func main() {
	// 文件服务器
	http.Handle("/", http.FileServer(http.Dir(".")))
	// HTTP 服务侦听在本机 8080 端口
	http.ListenAndServe(":8080", nil)
}
