package main

import (
	"net/http"
	"github.com/grellyd/grellyddotcom/handlers"
)

func routerSetup() {
	http.HandleFunc("/about", handlers.MakeHandler(handlers.StaticHandler))
	http.HandleFunc("/status", handlers.MakeHandler(handlers.StaticHandler))
	http.HandleFunc("/quote", handlers.MakeHandler(handlers.StaticHandler))
    
	// Dynamic page routing
	http.HandleFunc("/blog", handlers.MakeHandler(handlers.BlogHandler))
	// Resource routing ala css
	http.HandleFunc("/css/", handlers.MakeHandler(handlers.ResourceHandler))
}

func serve() {
	http.ListenAndServe(":3000", nil)
}

func main() {
	routerSetup()
	serve()
}
