package main

import (
	"fmt"
	"os" //我们要用到os包中的env
)
func main() {
	var VERSION string
	os.Setenv("VERSION","v1.0.0")
	VERSION = os.Getenv("VERSION")
	fmt.Println(VERSION)
}
