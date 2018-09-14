package main

import (
	"net/http"
	"github.com/grellyd/grellyddotcom/handlers"
)

func serve() {
	http.ListenAndServe(":3000", nil)
}

func main() {
	registerRoutes()
	serve()
}

func registerRoutes() {
	handlers.RouterSetup()
}
