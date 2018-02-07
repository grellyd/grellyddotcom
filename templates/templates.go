package templates

import (
	"html/template"
)

var All = template.Must(template.ParseFiles("templates/status.html", "templates/edit.html", "templates/view.html", "templates/index.html"))

