package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/grellyd/grellyddotcom/pages"
	"github.com/grellyd/filelogging/globallogger"
)

// File handler for any file 
func File(w http.ResponseWriter, r *http.Request) {
	globallogger.Debug(fmt.Sprintf("Handling File\n"))
	sections, title, pending, err := decomposeURL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle file: %s", err.Error()), http.StatusInternalServerError)
	}
	err = pages.CheckExistence(sections, title, pending)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to handle file: %s", err.Error()), http.StatusInternalServerError)
	} else {
		filepath := "public"
		for _, section := range(sections) {
			filepath = fmt.Sprintf("%s/%s", filepath, section)
		}
		filepath = fmt.Sprintf("%s/%s.%s", filepath, title, pending)
		globallogger.Debug(fmt.Sprintf("Serving '%s'", filepath))
		http.ServeFile(w, r, filepath)
	}
}

// decomponseURL breaks a URL down into its sections and title for hugo's routing.
// / ->                         ['','']
// /resume ->                   ['', 'resume']
// /resume/ ->                  ['', 'resume', '']
// /blog/post ->                ['', 'blog', 'post']
// /blog/post/ ->               ['', 'blog', 'post', '']
// /css/grellyd.com ->          ['', 'css'   , 'grellyd.com']
// /images/xmas/2018/wct.jpg -> ['', 'images', 'xmas', '2018', 'wct.jpg']
func decomposeURL(url string) (sections []string, title string, pending pages.PageEnding, err error) {
	globallogger.Debug(fmt.Sprintf("decomposing '%s'", url))
	trimmedURL := strings.TrimRight(url, "/")
	components := strings.Split(trimmedURL, "/")
	globallogger.Debug(fmt.Sprintf("decomposed to '%v' of len '%d'", components, len(components)))
	
	if strings.Contains(components[len(components) - 1], ".") {
		// is a direct file with title and type
		sections = components[1:len(components) - 1]
		fileDetails := strings.Split(components[len(components) - 1], ".")
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
