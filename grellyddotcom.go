package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/grellyd/grellyddotcom/handlers"
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
	http.HandleFunc("/", handlers.Static)
	http.HandleFunc("/blog/", handlers.Blog)
	http.HandleFunc("/css/", handlers.CSS)
	http.HandleFunc("/images/", handlers.Images)
	return nil
}

// Top Level Err Handle
func checkError(err error) {
	if err != nil {
		globallogger.Fatal(err.Error())
		os.Exit(1)
	}
}
