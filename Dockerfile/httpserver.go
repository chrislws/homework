package main

import (
	"io"
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, r*http.Request) {
	io.WriteString(w, "200")}


func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
