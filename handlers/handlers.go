package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/grellyd/filelogging/globallogger"
	"github.com/grellyd/grellyddotcom/pages"
)

func QRGen(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.RequestURI: %v\n", r.RequestURI)

	path := "../pages/templates/qrgen.tmpl"

	b, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to readfile %s: %s\n", path, err.Error()), http.StatusInternalServerError)
	}

	t, err := template.New("qrgen").Parse(string(b))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse %s: %s\n", string(b), err.Error()), http.StatusInternalServerError)
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to execute '%s' %s\n", string(b), err.Error()), http.StatusInternalServerError)
	}

}

// File handler for any file
func File(w http.ResponseWriter, r *http.Request) {
	globallogger.Debug(fmt.Sprintf("Handling File\n"))
	sections, title, pending, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if err = pages.CheckExistence(sections, title, pending); err != nil {
		http.Error(w, fmt.Sprintf("unable to handle file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	filepath := "public"
	for _, section := range sections {
		filepath = fmt.Sprintf("%s/%s", filepath, section)
	}

	filepath = fmt.Sprintf("%s/%s.%s", filepath, title, pending)
	globallogger.Debug(fmt.Sprintf("Serving '%s'", filepath))
	http.ServeFile(w, r, filepath)
}

// decomponseURL breaks a URL down into its sections and title for hugo's routing.
// / ->                         [”,”]
// /resume ->                   [”, 'resume']
// /resume/ ->                  [”, 'resume', ”]
// /blog/post ->                [”, 'blog', 'post']
// /blog/post/ ->               [”, 'blog', 'post', ”]
// /css/grellyd.com ->          [”, 'css'   , 'grellyd.com']
// /images/xmas/2018/wct.jpg -> [”, 'images', 'xmas', '2018', 'wct.jpg']
// /favicon.ico ->              [”, 'favicon.ico']
func decomposeURL(url string) (sections []string, title string, pending pages.PageEnding, err error) {
	globallogger.Debug(fmt.Sprintf("decomposing '%s'", url))
	downcasedURL := strings.ToLower(url)
	trimmedURL := strings.TrimRight(downcasedURL, "/")
	components := strings.Split(trimmedURL, "/")
	globallogger.Debug(fmt.Sprintf("decomposed to '%v' of len '%d'", components, len(components)))

	if strings.Contains(components[len(components)-1], ".") {
		// is a direct file with title and type
		sections = components[1 : len(components)-1]
		fileDetails := strings.Split(components[len(components)-1], ".")
		title = fileDetails[0]
		pending, err := pages.MatchPageEnding(fileDetails[1])
		if err != nil {
			err = fmt.Errorf("unable to decomposeURL: %s", err.Error())
		}
		globallogger.Debug(fmt.Sprintf("sections: %v; title: %s; pending: %s; err: %v", sections, title, pending, err))
		return sections, title, pending, err
	}
	// is a page browser
	return components[1:], "index", pages.HTML, err
}
