package main

import (
	"fmt"
	"log"
	"net/http"
)

type BaseHander struct {
}

func (hander *BaseHander) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path=>", req.URL.Path)
	fmt.Println(req.Body)
	fmt.Println("url param a =>", req.URL.Query().Get("a"))
	resp.Write([]byte("hello world"))
}

func main() {
	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/hello", helloHandler)
	server := &http.Server{Addr: "0.0.0.0:9999", Handler: &BaseHander{}}
	if err := server.ListenAndServe(); nil != err {
		log.Fatalf("listen and serve failed: " + err.Error())
	}
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	// 解析参数默认是不会解析
	req.ParseForm()
	fmt.Println("path", req.URL.Path)
	w.Write([]byte("Hello Golang"))
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
