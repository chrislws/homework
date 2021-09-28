package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//使用Get方法获取服务器响应包数据
	resp, err := http.Get("http://localhost:8080")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// 获取服务器端读到的数据
	fmt.Println("Request = ", resp.Request)
	fmt.Println("StatusCode = ", resp.StatusCode) // 状态码
	//读取body内的内容
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))

}