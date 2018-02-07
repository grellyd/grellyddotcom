package main

import (
	"net/http"
	"grellyddotcom/handlers"
)


// takes a list of templates
func router_setup() {
	http.HandleFunc("/", handlers.MakeHandler(handlers.IndexHandler))
    http.HandleFunc("/blog/", handlers.MakeHandler(handlers.BlogHandler))
	http.HandleFunc("/status/", handlers.MakeHandler(handlers.StatusHandler))
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))
}

func serve() {
	http.ListenAndServe(":3000", nil)
}

func main() {
	router_setup()
	serve()
}
