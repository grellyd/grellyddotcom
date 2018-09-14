package main

import (
	"net/http"
	"github.com/grellyd/grellyddotcom/handlers"
	"github.com/grellyd/grellyddotcom/router"
)

func serve() {
	http.ListenAndServe(":3000", nil)
}

func main() {
	registerRoutes()
	serve()
}

func registerRoutes() {
	err := router.Register("/", "^/$", router.NOARGS)
	handlers.
	handlers.RouterSetup()
}
