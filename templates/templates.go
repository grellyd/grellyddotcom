package templates

import (
	"html/template"
)

var All = template.Must(template.ParseFiles("status.html", "edit.html", "view.html", "index.html"))

