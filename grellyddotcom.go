package main

import (
	"fmt"
	"net/http"
	"grellyddotcom/handlers"
	"grellyddotcom/templates"
	"grellyddotcom/pages"
)


// takes a list of templates
func router_setup() {
	http.HandleFunc("/", makeHandler(indexHandler))
    http.HandleFunc("/blog/", makeHandler(blogHandler))
	http.HandleFunc("/status/", makeHandler(statusHandler))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
}

func serve() {
	http.ListenAndServe(":3000", nil)
}

func main() {
	router_setup()
	serve()
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
