package main

import (
	"net/http"
	"github.com/grellyd/grellyddotcom/handlers"
)

func serve() {
	http.ListenAndServe(":3000", nil)
}

func main() {
	handlers.RouterSetup()
	serve()
}
