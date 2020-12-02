package router

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/grellyd/filelogging/globallogger"
)

// Route is a route routed by the router
type Route struct {
	path    string
	reg     *regexp.Regexp
	handler http.Handler
}

type Router struct {
	Mux    *http.ServeMux
	routes []Route
}

func NewRouter() *Router {
	return &Router{
		Mux:    http.NewServeMux(),
		routes: []Route{},
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Mux.ServeHTTP(w, req)
}

// Register takes a url path and a valid regexp for the route. Returns an error if there is a problem
func (r *Router) Register(path string, validRegexp string, handler func(http.ResponseWriter, *http.Request)) error {
	rex, err := regexp.Compile(validRegexp)
	if err != nil {
		return fmt.Errorf("unable to compile '%s' as a regexp for '%s': %s", validRegexp, path, err.Error())
	}

	outerHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		globallogger.Info(fmt.Sprintf("Handling route: %s\n", req.URL.Path))
		handler(w, req)
	})

	r.Mux.HandleFunc(path, outerHandler)

	route := Route{
		path:    path,
		reg:     rex,
		handler: http.HandlerFunc(handler),
	}

	r.routes = append(r.routes, route)

	return nil
}
