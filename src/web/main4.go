package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"web/mux"
)

type Auth struct {
	username string `json:"username"`
	password string `json:"password"`
}

type BaseHander struct {
}

func (hander *BaseHander) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		req.ParseForm()
		fmt.Println("method:", req.Method) //获取请求的方法
		fmt.Println("username", req.Form["username"])
		fmt.Println("password", req.Form["password"])
		for k, v := range req.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if req.Method == "POST" {
		var auth Auth
		if err := json.NewDecoder(req.Body).Decode(&auth); err != nil {
			log.Fatal(err)
		}
		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()
		fmt.Println(auth)

		//未知类型的推荐处理方法
		var f interface{}
		json.Unmarshal(result, &f)
		m := f.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}
	resp.Write([]byte("hello world"))
}

var (
	port string
)

func main() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.Parse()

	router := mux.NewMuxHandler()
	router.Handle("/hello/golang/", &BaseHander{})
	router.HandleFunc("/hello/world", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))
	})
	http.ListenAndServe(port, router)

}
