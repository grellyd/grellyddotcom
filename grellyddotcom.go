package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/grellyd/filelogging/globallogger"
	"github.com/grellyd/filelogging/state"
	"github.com/grellyd/grellyddotcom/handlers"
	"github.com/grellyd/grellyddotcom/router"
)

func main() {
	err := setup()
	checkError(err)
	globallogger.Info("Setup Complete")
	http.ListenAndServe(":8080", nil)
}

func setup() (err error) {
	err = globallogger.NewGlobalLogger("grellyd.com Server", state.NORMAL)
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
	err = router.Register("/", "(^/$)|(^/(status|about|quote|xmas)$)", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register static: %s", err.Error())
	}
	err = router.Register("/blog/", "^/blog/([a-zA-Z0-9]*)$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register blog: %s", err.Error())
	}
	err = router.Register("/css/", "^/css/([a-zA-Z0-9_]*).css$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register css: %s", err.Error())
	}
	err = router.Register("/js/", "^/js/([a-zA-Z0-9_]*).js$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register js: %s", err.Error())
	}
	err = router.Register("/images/", "^/images/(([a-zA-Z0-9_]*)/)*.(jpg|png)$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register images: %s", err.Error())
	}
	err = router.Register("/videos/", "^/videos/(([a-zA-Z0-9_]*)/)*.mp4$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register videos: %s", err.Error())
	}
	err = router.Register("/files/", "^/files/(([a-zA-Z0-9_]*)/)*.jpg$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register files: %s", err.Error())
	}
	err = router.Register("/xmas/", "^/xmas/([a-zA-Z0-9_]*).html$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register xmas: %s", err.Error())
	}
	err = router.Register("/email/", "^/email/([a-zA-Z0-9_]*).html$", handlers.File)
	if err != nil {
		return fmt.Errorf("unable to register email: %s", err.Error())
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
