package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
@Description: 简单的一个web服务器
@Date: 2/18/2021
@Author: lichang
*/

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ", "))
	}
	_, _ = fmt.Fprintf(w, "base lichang")
}

func main() {

	http.HandleFunc("/", sayHelloName)
	/*  func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
		 	DefaultServeMux.HandleFunc(pattern, handler)
	    }
		注意到这里是用了DefaultServeMux这个对象，为其设置了handler
	*/

	err := http.ListenAndServe(":9090", nil)
	/*
		...
		...

		func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
			handler := sh.srv.Handler
			if handler == nil {
				handler = DefaultServeMux
			}
			if req.RequestURI == "*" && req.Method == "OPTIONS" {
				handler = globalOptionsHandler{}
			}
			handler.ServeHTTP(rw, req)
		}
		如果在ListenAndServe里设置第二个参数为nil，那么最后的handler也会是DefaultServeMux
	*/

	// 如果我们设置ListenAndServe第二个参数设置为自定义的Handler，即可将http请求交给自定义Handler处理

	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
