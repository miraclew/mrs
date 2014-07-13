package main

import (
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

func httpServe(listener net.Listener) {
	log.Printf("HTTP: listening on %s", listener.Addr().String())

	handler := http.NewServeMux()

	handler.Handle("/user", restful.NewRestfulApiHandler(new(api.UserController)))
	handler.Handle("/token", restful.NewRestfulApiHandler(new(api.TokenController)))
	handler.Handle("/match", restful.NewRestfulApiHandler(new(api.MatchController)))

	handler.HandleFunc("/", root)

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
