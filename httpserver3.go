package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//注册回调函数
	http.HandleFunc("/hello", hellohello)
	log.Println("start")
	//绑定tcp监听地址，并开始接受请求，然后调用服务端处理程序来处理传入的连接请求
	//参数1为addr即监听地址；参数2表示服务端处理程序，通常为nil
	//当参数2为nil时，服务端调用http.DefaultServeMux进行处理
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hellohello(w http.ResponseWriter, r *http.Request) {
	//ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段
	r.ParseForm()
	fmt.Println(r.Form)
	//遍历打印解析结果
	for key, value := range r.Form {
		fmt.Println(key, value)

	}
	w.Write([]byte("hello"))
}