package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"github.com/grellyd/grellyddotcom/pages"
	//"github.com/grellyd/grellyddotcom/layouts"
)

var rootPath = regexp.MustCompile("^/$")
// TODO: handle ending slash
var staticPath = regexp.MustCompile("^/(status|about|quote)$")
var blogPath = regexp.MustCompile("^/blog/([a-zA-Z0-9]*)$")

// MakeHandler creates a function to call when handling a route.
func MakeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path
		switch {
		case rootPath.MatchString(requestPath):
			fn(w, r, "index")
			return
		case staticPath.MatchString(requestPath):
			staticTitle := strings.Split(requestPath, "/")[1]
			fn(w, r, staticTitle)
			return
		case blogPath.MatchString(requestPath):
			blogTitle := strings.Split(requestPath, "/")[2]
			fn(w, r, blogTitle)
		default:
			fmt.Printf("'%s' is an invalid path\n", r.URL.Path)
			http.NotFound(w, r)
			return
		}
	}
}

// StaticHandler handles any static page
func StaticHandler(w http.ResponseWriter, r *http.Request, title string) {
	_, err := pages.Load(pages.STATIC, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("public/static/%s.html", title))
}

// BlogHandler manages selecting the correct blog page
func BlogHandler(w http.ResponseWriter, r *http.Request, title string) {
	// TODO: handle multi
	_, err := pages.Load(pages.BLOG, title)
	if err != nil {
		http.Redirect(w, r, "/blog/" + title, http.StatusFound)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("public/blog/%s.html", title))
}
