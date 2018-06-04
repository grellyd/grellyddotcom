package content

import (
	"html/template"
)

// TODO: Table static special types, like about, index, contact, etc.

var wikiFiles = []string{"layouts/wiki/status.html", "layouts/wiki/edit.html", "layouts/wiki/view.html"}
var defaultFiles = []string{"layouts/_default/single.html", "layouts/_default/list.html"}

// Wiki collection
var Wiki = template.Must(template.ParseFiles(wikiFiles...))

// Default collection
var Default = template.Must(template.ParseFiles(defaultFiles...))

// All collection
var All = template.Must(template.ParseFiles(append(defaultFiles, wikiFiles...)...))
