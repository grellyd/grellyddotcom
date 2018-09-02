package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"github.com/grellyd/grellyddotcom/pages"
)

// TODO: Overall refactoring needed. Currently every resource is loaded twice per request.

var rootPath = regexp.MustCompile("^/$")
// TODO: handle ending slash
var staticPath = regexp.MustCompile("^/(status|about|quote)$")
var blogPath = regexp.MustCompile("^/blog/([a-zA-Z0-9]*)$")
var resourcePath = regexp.MustCompile("^/(css|images)/([a-zA-Z0-9_]*).(css|jpg)$")

// RouterSetup sets up the http routes available to the webapp
func RouterSetup() {
	// Static pages
	http.HandleFunc("/", makeHandler(staticHandler))
	// Dynamic page routing
	http.HandleFunc("/blog", makeHandler(blogHandler))
	// Resource routing
	http.HandleFunc("/css/", makeHandler(cssHandler))
	http.HandleFunc("/images/", makeHandler(imagesHandler))
}

// makeHandler creates a function to call when handling a route, and passes the correct arguments
// TODO: Refactor, as currently this is little more than a glorified argument passer
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
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
		case resourcePath.MatchString(requestPath):
			resourceTitle := strings.Split(requestPath, "/")[2]
			fn(w, r, resourceTitle)
		default:
			fmt.Printf("'%s' is an invalid path\n", r.URL.Path)
			http.NotFound(w, r)
			return
		}
	}
}

// StaticHandler handles any static page
func staticHandler(w http.ResponseWriter, r *http.Request, title string) {
	err := pages.CheckExistence(pages.STATIC, title, pages.HTML)
	if err != nil {
		http.ServeFile(w, r, fmt.Sprintf("public/%s.%s", title, pages.HTML))
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// BlogHandler manages selecting the correct blog page
func blogHandler(w http.ResponseWriter, r *http.Request, title string) {
	// TODO: handle multi
	_, err := pages.Load(pages.BLOG, title, pages.HTML)
	if err != nil {
		http.Redirect(w, r, "/blog/" + title, http.StatusFound)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("public/blog/%s.%s", title, pages.HTML))
}

// ImageHandler manages serving the correct resource file
func imagesHandler(w http.ResponseWriter, r *http.Request, title string) {
	resourceName := strings.Split(title, ".")[0]
	_, err := pages.Load(pages.IMAGESRESOURCE, resourceName, pages.JPG)
	if err != nil {
		fmt.Println(err.Error())
		// TODO: change from do nothing on resource not found. 
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("public/images/%s.%s", resourceName, pages.JPG))
}

// CSSHandler manages serving the correct resource file
func cssHandler(w http.ResponseWriter, r *http.Request, title string) {
	resourceName := strings.Split(title, ".")[0]
	_, err := pages.Load(pages.CSSRESOURCE, resourceName, pages.CSS)
	if err != nil {
		fmt.Println(err.Error())
		// TODO: change from do nothing on resource not found. 
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("public/css/%s.%s", resourceName, pages.CSS))
}
