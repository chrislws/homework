package main

import (
	"fmt"
	"net/http"
)

//w, 给客户端回复数据
//r, 读取客户端发送的数据
func HandConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.RemoteAddr = ", r.RemoteAddr)
	w.Write([]byte("hello go")) //给客户端回复数据
}

func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/", HandConn)

	//监听绑定
	http.ListenAndServe(":8080", nil)
}