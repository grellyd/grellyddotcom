package handlers

import (
	"fmt"
	"html/template"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/grellyd/filelogging/globallogger"
	"github.com/grellyd/grellyddotcom/pages"
	"github.com/grellyd/qrgen/qrlib"
)

func QRGen(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintln("Wrong method"), http.StatusBadRequest)
		return
	}
	fmt.Printf("r.RequestURI: %v\n", r.RequestURI)

	path := "./templates/qrgen.tmpl"

	globallogger.Info("in qrgen")

	b, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to readfile %s: %s\n", path, err.Error()), http.StatusInternalServerError)
		return
	}

	t, err := template.New("qrgen").Parse(string(b))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse %s: %s\n", string(b), err.Error()), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to execute '%s' %s\n", string(b), err.Error()), http.StatusInternalServerError)
		return
	}
}

func generateQRCode(link string, scaleFactor int) (*image.Gray, error) {
	qr, err := qrlib.Generate(link, qrlib.ECLH)
	if err != nil {
		return nil, fmt.Errorf("unable to generate: %s", err.Error())
	}

	i, err := qrlib.BuildImage(qr, scaleFactor)
	if err != nil {
		return nil, fmt.Errorf("unable to build an output image: %s", err.Error())
	}

	return i, nil
}

func QRCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintln("Wrong method: Got %s; expected %s", r.Method, http.MethodPost), http.StatusBadRequest)
		return
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	globallogger.Info("in qrcode")
	fmt.Printf("r.RequestURI: %v\n", r.RequestURI)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to ReadAll: %s\n", err.Error()), http.StatusInternalServerError)
		return
	}
	globallogger.Info(fmt.Sprintf("body: %s\n", string(body)))

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("failed to parseForm: %s\n", err.Error()), http.StatusInternalServerError)
		return
	}

	for k, v := range r.Form {
		globallogger.Info(fmt.Sprintf("%s: %s\n", k, v))
	}

	for k, v := range r.PostForm {
		globallogger.Info(fmt.Sprintf("%s: %s\n", k, v))
	}

	link := r.PostFormValue("link")
	globallogger.Info(fmt.Sprintf("link: %s\n", link))

	scaleFactor := 50

	i, err := generateQRCode(link, scaleFactor)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to generateQRCode for %s: %s\n", link, err.Error()), http.StatusInternalServerError)
		return
	}

	uuid := uuid.NewString()
	path := fmt.Sprintf("/images/qrcodes/%s.png", uuid)

	f, err := os.Create(fmt.Sprintf("/var/http/public/%s", path))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create file: %s\n", err.Error()), http.StatusInternalServerError)
		return
	}

	defer f.Close()
	if err = png.Encode(f, i); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode png: %s\n", err.Error()), http.StatusInternalServerError)
		return
	}

	b, err := os.ReadFile("./templates/qrcode.tmpl")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to readfile %s: %s\n", path, err.Error()), http.StatusInternalServerError)
		return
	}

	t, err := template.New("qrgen").Parse(string(b))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse %s: %s\n", string(b), err.Error()), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, map[string]any{"image": path})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to execute '%s' %s\n", string(b), err.Error()), http.StatusInternalServerError)
		return
	}

}

// File handler for reading any file
func File(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintln("Wrong method"), http.StatusBadRequest)
		return
	}
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
