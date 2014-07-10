package main

import (
	"fmt"
	"github.com/miraclew/mrs/api"
	"github.com/miraclew/restful"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func init() {
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello root page!") //这个写入到w的是输出到客户端的
}

func httpServe(listener net.Listener) {
	log.Printf("HTTP: listening on %s", listener.Addr().String())

	handler := http.NewServeMux()

	handler.Handle("/v1/user", restful.NewRestfulApiHandler(new(api.UserController)))
	handler.Handle("/v1/token", restful.NewRestfulApiHandler(new(api.TokenController)))

	handler.HandleFunc("/", sayhelloName)

	server := &http.Server{
		Handler: handler,
	}

	err := server.Serve(listener)
	// theres no direct way to detect this error because it is not exposed
	if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		log.Printf("ERROR: http.Serve() - %s", err.Error())
	}

	log.Printf("HTTP: closing %s", listener.Addr().String())
}

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "welcome to mrs server.")
}
