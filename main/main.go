package main

import (
	"net/http"
	_ "net/http/pprof"
	"test/test1"
)

func test1Handler(w http.ResponseWriter, r *http.Request) {
	test1.WritePng(w)
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
	data := []byte("<html><head><title>Hello world</title></head><body><h1>Hello World!</h1></body></html>")
	w.Write(data)
}

func main() {
	http.HandleFunc("/test1", test1Handler)
	http.HandleFunc("/test2", test2Handler)
	http.ListenAndServe(":9090", nil)
}
