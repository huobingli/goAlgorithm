package main
import (
    "net/http"
)
func main() {
	// 文件服务器
	http.Handle("/", http.FileServer(http.Dir(".")))
	// HTTP 服务侦听在本机 8080 端口
    http.ListenAndServe(":8080", nil)
}