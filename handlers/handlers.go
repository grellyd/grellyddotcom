package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/grellyd/grellyddotcom/pages"
	"github.com/grellyd/filelogging/globallogger"
)

// TODO: Remove duplication

// Static handles any static page
func Static(w http.ResponseWriter, r *http.Request) {
	section, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle static page: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(section, title, pages.HTML)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle static page: %s", err.Error()), http.StatusInternalServerError)
	} else {
		filepath := ""
		if len(section) == 0 {
			filepath = fmt.Sprintf("public/%s.%s", title, pages.HTML)
		} else {
			filepath = fmt.Sprintf("public/%s/%s.%s", section, title, pages.HTML)
		}
		globallogger.Debug(fmt.Sprintf("Serving '%s'", filepath))
		http.ServeFile(w, r, filepath)
	}
}

// Blog manages selecting the correct blog page
func Blog(w http.ResponseWriter, r *http.Request) {
	section, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle blog page: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(section, title, pages.HTML)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle blog page: %s", err.Error()), http.StatusInternalServerError)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("public/blog/%s.%s", title, pages.HTML))
	}
}

// Images manages serving the correct resource file
func Images(w http.ResponseWriter, r *http.Request) {
	section, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle image: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(section, title, pages.JPG)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle image: %s", err.Error()), http.StatusInternalServerError)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("public/images/%s.%s", title, pages.JPG))
	}
}

// CSS manages serving the correct resource file
func CSS(w http.ResponseWriter, r *http.Request) {
	section, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle image: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(section, title, pages.CSS)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle image: %s", err.Error()), http.StatusInternalServerError)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("public/css/%s.%s", title, pages.CSS))
	}
}

// decomponseURL breaks a URL down into its section and title for hugo's routing.
// / ->              ['','']
// /resume ->        ['', 'resume']
// /resume/ ->       ['', 'resume', '']
// /resume/index ->  ['', 'resume', 'index']
// /resume/index/ -> ['', 'resume', 'index', '']
func decomposeURL(url string) (section string, title string, err error) {
	globallogger.Debug(fmt.Sprintf("decomposing '%s'", url))
	trimmedURL := strings.TrimRight(url, "/")
	components := strings.Split(trimmedURL, "/")
	globallogger.Debug(fmt.Sprintf("decomposed to '%v' of len '%d'", components, len(components)))
	switch len(components) {
	case 1: 
		// root url
		section = ""
		title = "index"
	case 2:
		// section header
		section = components[1]
		title = "index"
	case 3:
		// section page
		section = components[1]
		title = strings.Split(components[2], ".")[0]
	default:
		err = fmt.Errorf("unhandled url structure")
	}
	return section, title, err
}
