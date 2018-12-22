package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/grellyd/grellyddotcom/handlers"
	"github.com/grellyd/grellyddotcom/router"
	"github.com/grellyd/filelogging/globallogger"
	"github.com/grellyd/filelogging/state"
)

func main() {
	err := setup()
	checkError(err)
	globallogger.Info("Setup Complete")
	// http.ListenAndServe(":3000", http.FileServer(http.Dir("public/")))
	http.ListenAndServe(":3000", nil)
}

func setup() (err error) {
	err = globallogger.NewGlobalLogger("grellyd.com Server", state.DEBUGGING)
	if err != nil {
		return fmt.Errorf("setup failed: %s", err.Error())
	}
	err = registerRoutes()
	if err != nil {
		return fmt.Errorf("setup failed: %s", err.Error())
	}
	return nil
}

func registerRoutes() (err error) {
	err = router.Register("/", "(^/$)|(^/(status|about|quote|xmas)$)", handlers.Static)
	if err != nil {
		return fmt.Errorf("unable to register static: %s", err.Error())
	}
	err = router.Register("/blog/", "^/blog/([a-zA-Z0-9]*)$", handlers.Blog)
	if err != nil {
		return fmt.Errorf("unable to register blog: %s", err.Error())
	}
	err = router.Register("/css/", "^/css/([a-zA-Z0-9_]*).css$", handlers.CSS)
	if err != nil {
		return fmt.Errorf("unable to register css: %s", err.Error())
	}
	err = router.Register("/images/", "^/images/([a-zA-Z0-9_]*).jpg$", handlers.Images)
	if err != nil {
		return fmt.Errorf("unable to register images: %s", err.Error())
	}
	err = router.Register("/files/", "^/files/([a-zA-Z0-9_]*).jpg$", handlers.Files)
	if err != nil {
		return fmt.Errorf("unable to register files: %s", err.Error())
	}
	err = router.Register("/xmas/", "^/xmas/([a-zA-Z0-9_]*).html$", handlers.Xmas)
	if err != nil {
		return fmt.Errorf("unable to register xmas: %s", err.Error())
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		// Top Level Err Handle
		globallogger.Fatal(err.Error())
		os.Exit(1)
	}
}
