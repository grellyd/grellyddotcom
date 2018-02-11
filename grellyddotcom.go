package main

import (
	"net/http"
	"grellyddotcom/handlers"
)


func router_setup() {
	http.HandleFunc("/", handlers.MakeHandler(handlers.IndexHandler))
    http.HandleFunc("/blog/", handlers.MakeHandler(handlers.BlogHandler))
	http.HandleFunc("/status/", handlers.MakeHandler(handlers.StatusHandler))
}

func serve() {
	http.ListenAndServe(":3000", nil)
}

func main() {
	router_setup()
	serve()
}
