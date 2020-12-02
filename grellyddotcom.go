package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/grellyd/filelogging/globallogger"
	"github.com/grellyd/filelogging/state"
	"github.com/grellyd/grellyddotcom/handlers"
	"github.com/grellyd/grellyddotcom/router"
)

func main() {
	r, certManager, err := setup()
	if err != nil {
		checkError(fmt.Errorf("failed to setup: %w", err))
	}
	globallogger.Info("Setup Complete")

	server := &http.Server{
		Addr: ":8443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12,
		},
		Handler: r,
	}

	go func() {
		globallogger.Info("Serving Challenges")
		err = http.ListenAndServe(":8080", certManager.HTTPHandler(nil))
		if err != nil {
			globallogger.Error(fmt.Errorf("failed to listen and serve :8080: %w", err).Error())
			fmt.Println(err.Error())
		}
	}()

	globallogger.Info("Serving TLS")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		checkError(fmt.Errorf("failed to ListenAndServe: %w", err))
	}

}

func setup() (*router.Router, *autocert.Manager, error) {
	err := globallogger.NewGlobalLogger("grellyd.com Server", state.NORMAL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to globallogger.NewGlobalLogger: %w", err)
	}
	r, err := registerRoutes()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to registerRoutes: %w", err)
	}

	certManager := autocert.Manager{
		Prompt:      autocert.AcceptTOS,
		HostPolicy:  autocert.HostWhitelist("grellyd.com"),
		Cache:       autocert.DirCache("certs"),
		RenewBefore: 24 * time.Hour,
	}

	return r, &certManager, nil
}

func registerRoutes() (*router.Router, error) {
	r := router.NewRouter()
	err := r.Register("/", "(^/$)|(^/(status|about|quote|xmas)$)", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register static: %w", err)
	}
	err = r.Register("/blog/", "^/blog/([a-zA-Z0-9]*)$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register blog: %w", err)
	}
	err = r.Register("/css/", "^/css/([a-zA-Z0-9_]*).css$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register css: %w", err)
	}
	err = r.Register("/js/", "^/js/([a-zA-Z0-9_]*).js$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register js: %w", err)
	}
	err = r.Register("/images/", "^/images/(([a-zA-Z0-9_]*)/)*.(jpg|png)$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register images: %w", err)
	}
	err = r.Register("/videos/", "^/videos/(([a-zA-Z0-9_]*)/)*.mp4$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register videos: %w", err)
	}
	err = r.Register("/files/", "^/files/(([a-zA-Z0-9_]*)/)*.jpg$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register files: %w", err)
	}
	err = r.Register("/xmas/", "^/xmas/([a-zA-Z0-9_]*).html$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register xmas: %w", err)
	}
	err = r.Register("/email/", "^/email/([a-zA-Z0-9_]*).html$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register email: %w", err)
	}
	err = r.Register("/certs/", "^/certs/([a-zA-Z0-9_]*).html$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register certs: %w", err)
	}

	return r, nil
}

func checkError(err error) {
	if err != nil {
		// Top Level Err Handle
		globallogger.Fatal(err.Error())
		os.Exit(1)
	}
}
