package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/acme/autocert"

	"github.com/grellyd/filelogging/globallogger"
	"github.com/grellyd/filelogging/state"
	"github.com/grellyd/grellyddotcom/config"
	"github.com/grellyd/grellyddotcom/router"
)

const (
	AddrHTTP  = ":80"
	AddrHTTPS = ":443"
)

func main() {
	err := run()
	if err != nil {
		globallogger.Fatal(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {

	config, err := config.NewConfig(os.Args)
	if err != nil {
		return errors.Wrap(err, "failed to NewConfig")
	}

	r, err := buildRouter()
	if err != nil {
		checkError(fmt.Errorf("failed to setup: %w", err))
	}

	server := &http.Server{
		Addr:    AddrHTTP,
		Handler: r,
	}

	if config.TLS {
		certManager, err := buildCertManager()
		if err != nil {
			return errors.Wrap(err, "failed to buildCertManager")
		}

		server.TLSConfig = &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12,
		}
		server.Addr = AddrHTTPS

		go func() {
			globallogger.Info("Serving Challenges")
			err = http.ListenAndServe(AddrHTTP, certManager.HTTPHandler(nil))
			if err != nil {
				globallogger.Error(fmt.Errorf("failed to listen and serve %s: %w", AddrHTTP, err).Error())
				fmt.Println(err.Error())
			}
		}()

		globallogger.Info(fmt.Sprintf("Serving on %s", server.Addr))
		if err = server.ListenAndServeTLS("", ""); err != nil {
			return errors.Wrap(err, "failed to ListenAndServeTLS")
		}

	} else {
		globallogger.Info(fmt.Sprintf("Serving on %s", server.Addr))
		if err = server.ListenAndServe(); err != nil {
			return errors.Wrap(err, "failed to ListenAndServe")
		}
	}

	return nil

}

func buildCertManager() (*autocert.Manager, error) {
	certManager := autocert.Manager{
		Prompt:      autocert.AcceptTOS,
		HostPolicy:  autocert.HostWhitelist("grellyd.com", "www.grellyd.com", "dev.grellyd.com"),
		Cache:       autocert.DirCache("certs"),
		RenewBefore: 24 * time.Hour,
	}

	return &certManager, nil
}

func buildRouter() (*router.Router, error) {
	err := globallogger.NewGlobalLogger("grellyd.com Server", state.NORMAL)
	if err != nil {
		return nil, fmt.Errorf("failed to globallogger.NewGlobalLogger: %w", err)
	}
	r, err := registerRoutes()
	if err != nil {
		return nil, fmt.Errorf("failed to registerRoutes: %w", err)
	}

	return r, nil
}

func registerRoutes() (*router.Router, error) {
	r := router.NewRouter()
	// err := r.Register("/qrgen", "^/qrgen$", handlers.QRGen)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register qrgen: %w", err)
	// }
	// err = r.Register("/", "(^/$)|(^/(status|about|quote|xmas)$)", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register static: %w", err)
	// }
	// err = r.Register("/writing/", "^/writing/([a-zA-Z0-9]*)$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register writing: %w", err)
	// }
	// err = r.Register("/css/", "^/css/([a-zA-Z0-9_]*).css$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register css: %w", err)
	// }
	// err = r.Register("/js/", "^/js/([a-zA-Z0-9_]*).js$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register js: %w", err)
	// }
	// err = r.Register("/images/", "^/images/(([a-zA-Z0-9_]*)/)*.(jpg|png)$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register images: %w", err)
	// }
	// err = r.Register("/videos/", "^/videos/(([a-zA-Z0-9_]*)/)*.mp4$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register videos: %w", err)
	// }
	// err = r.Register("/files/", "^/files/(([a-zA-Z0-9_]*)/)*.jpg$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register files: %w", err)
	// }
	// err = r.Register("/xmas/", "^/xmas/([a-zA-Z0-9_]*).html$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register xmas: %w", err)
	// }
	// err = r.Register("/email/", "^/email/([a-zA-Z0-9_]*).html$", handlers.File)
	// if err != nil {
	// return nil, fmt.Errorf("unable to register email: %w", err)
	// }
	err = r.Register("/certs/", "^/certs/([a-zA-Z0-9_]*).html$", handlers.File)
	if err != nil {
		return nil, fmt.Errorf("unable to register certs: %w", err)
	}

	return r, nil
}

func checkError(err error) {
	if err != nil {
		// Top Level Err Handle
		os.Exit(1)
	}
}
