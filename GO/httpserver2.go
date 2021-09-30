package main

import (
	"fmt"
	"net/http"
	"os" //我们要用到os包中的env
)
func main() {
	var VERSION string
	os.Setenv("VERSION","v1.0.0")
	VERSION = os.Getenv("VERSION")
	fmt.Println(VERSION)
}

func VERSION(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("v1.0.0"))
}