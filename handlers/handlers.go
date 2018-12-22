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
	sections, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle static page: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(sections, title, pages.HTML)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle static page: %s", err.Error()), http.StatusInternalServerError)
	} else {
		filepath := ""
		if len(sections) == 0 {
			filepath = fmt.Sprintf("public/%s.%s", title, pages.HTML)
		} else {
			filepath = fmt.Sprintf("public/%s/%s.%s", sections, title, pages.HTML)
		}
		globallogger.Debug(fmt.Sprintf("Serving '%s'", filepath))
		http.ServeFile(w, r, filepath)
	}
}

// Blog manages selecting the correct blog page
func Blog(w http.ResponseWriter, r *http.Request) {
	sections, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle blog page: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(sections, title, pages.HTML)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle blog page: %s", err.Error()), http.StatusInternalServerError)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("public/blog/%s.%s", title, pages.HTML))
	}
}

// Xmas manages selecting the correct xmas page
func Xmas(w http.ResponseWriter, r *http.Request) {
	globallogger.Debug(fmt.Sprintf("Handling Xmas Page\n"))
	sections, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle xmas page: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(sections, title, pages.HTML)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle xmas page: %s", err.Error()), http.StatusInternalServerError)
	} else {
		filepath := "public"
		for _, section := range(sections) {
			filepath = fmt.Sprintf("%s/%s", filepath, section)
		}
		filepath = fmt.Sprintf("%s/%s.%s", filepath, title, pages.HTML)
		globallogger.Debug(fmt.Sprintf("Serving '%s'", filepath))
		http.ServeFile(w, r, filepath)
	}
}

// Images manages serving the correct resource file
func Images(w http.ResponseWriter, r *http.Request) {
	sections, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle image: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(sections, title, pages.JPG)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle image: %s", err.Error()), http.StatusInternalServerError)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("public/images/%s.%s", title, pages.JPG))
	}
}

// Files manages serving the correct resource file
func Files(w http.ResponseWriter, r *http.Request) {
	sections, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle file: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(sections, title, pages.PDF)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle file: %s", err.Error()), http.StatusInternalServerError)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("public/files/%s.%s", title, pages.PDF))
	}
}

// CSS manages serving the correct resource file
func CSS(w http.ResponseWriter, r *http.Request) {
	sections, title, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle css: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(sections, title, pages.CSS)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle css: %s", err.Error()), http.StatusInternalServerError)
	} else {
		http.ServeFile(w, r, fmt.Sprintf("public/css/%s.%s", title, pages.CSS))
	}
}

// decomponseURL breaks a URL down into its sections and title for hugo's routing.
// / ->              ['','']
// /resume ->        ['', 'resume']
// /resume/ ->       ['', 'resume', '']
// /resume/index ->  ['', 'resume', 'index']
// /resume/index/ -> ['', 'resume', 'index', '']
func decomposeURL(url string) (sections []string, title string, err error) {
	globallogger.Debug(fmt.Sprintf("decomposing '%s'", url))
	trimmedURL := strings.TrimRight(url, "/")
	components := strings.Split(trimmedURL, "/")
	globallogger.Debug(fmt.Sprintf("decomposed to '%v' of len '%d'", components, len(components)))
	switch len(components) {
	case 1: 
		// root url
		sections = []string{}
		title = "index"
	case 2:
		// sections header
		sections = []string{components[1]}
		title = "index"
	case 3:
		if strings.Contains(components[2], "."){
			// sections page
			sections = []string{components[1]}
			title = strings.Split(components[2], ".")[0]
		} else {
			// TODO: bandaid for three levels, not n
			// multiple setion page
			sections = components[1:]
			title = "index"
		}
	default:
		err = fmt.Errorf("unhandled url structure")
	}
	return sections, title, err
}
