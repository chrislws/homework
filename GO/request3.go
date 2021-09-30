package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, _ := http.Get("http://127.0.0.1:8080")
	fmt.Println("RemoteAddr = ", response.Request)
	fmt.Println("StatusCode = ", response.StatusCode) // 状态码
}