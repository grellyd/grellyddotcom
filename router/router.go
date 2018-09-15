package router

import (
	"fmt"
	"net/http"
	"regexp"
)

// Route is a route routed by the router
type Route struct{
	path string;
	reg *regexp.Regexp;
	handler http.Handler
}

// Routes are all the registered routes
var Routes []Route

// Register takes a url path and a valid regexp for the route. Returns an error if there is a problem
func Register(path string, validRegexp string, handler func (http.ResponseWriter, *http.Request)) error {
	r, err := regexp.Compile(validRegexp)
	if err != nil {
		return fmt.Errorf("unable to compile '%s' as a regexp for '%s': %s", validRegexp, path, err.Error())
	}
	http.HandleFunc(path, handler)
	route := Route{
		path: path,
		reg: r,
		handler: http.HandlerFunc(handler),
	}
	Routes = append(Routes, route)
	return nil
}
