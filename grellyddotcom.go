package main

import (
	"fmt"
	"regexp"
	"net/http"
	"io/ioutil"
	"html/template"
)

type Page struct {
	Title string
	Body  []byte
}

// takes a list of templates
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view|status)/([a-zA-Z0-9]*)$")

func router_setup() {
	// http.HandleFunc("/", handler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
    http.HandleFunc("/blog/", makeHandler(blogHandler))
	http.HandleFunc("/status/", makeHandler(statusHandler))
}

func serve() {
	http.ListenAndServe(":80", nil)
}

func main() {
	router_setup()
	serve()
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func renderTemplate(w http.ResponseWriter, tmplt string, p *Page) {
	err := templates.ExecuteTemplate(w, tmplt+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Fprintf(w, "OK")
}


func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// will become edit blog post
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
	renderTemplate(w, "edit", p)
}

func blogHandler(w http.ResponseWriter, r *http.Request, title string) {
}



// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("Invalid Page Title")
// 	}
// 	return m[2], nil
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }
// 
// func old_main() {
// 	p1 := &Page{Title: "IndexPage", Body: []byte("Welcome to GrellydDotCom!")}
// 	p1.save()
// 	p2, err := loadPage("IndexPage")
// 	if err == nil {
// 	    fmt.Println(string(p2.Body))
// 	}
// }
