package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"../pages"
	"../templates"
)

var rootPath = regexp.MustCompile("^/$")
var validPath = regexp.MustCompile("^/(blog|status|files)/([a-zA-Z0-9]*)$")

func MakeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Refactor
		path := r.URL.Path
		rootSubStrings := rootPath.FindStringSubmatch(path)
		otherSubStrings := validPath.FindStringSubmatch(path)
		if rootSubStrings == nil && otherSubStrings == nil {
			fmt.Printf("'%s' is an invalid path\n", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		if rootSubStrings != nil {
			fmt.Printf("Root substrings: %v\n", rootSubStrings)
			fn(w, r, "index")
			return
		} else {
			fmt.Printf("Other substrings: %v\n", otherSubStrings)
			fn(w, r, otherSubStrings[2])
			return
		}
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage("status")
	if err != nil {
		http.Redirect(w, r, "/status/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "status", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &pages.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func EditBlogHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
    if err != nil {
        p = &pages.Page{Title: title}
    }
	renderTemplate(w, "edit", p)
}

func BlogHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/blog/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "blog", p)
}

func IndexHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, "index", nil)
}

func renderTemplate(w http.ResponseWriter, tmplt string, p *pages.Page) {
	err := templates.All.ExecuteTemplate(w, tmplt+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
