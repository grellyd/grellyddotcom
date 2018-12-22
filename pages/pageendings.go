package pages

import "fmt"

// PageEnding is the type of a page indicating how the page should be treated
type PageEnding string 

const (
	// HTML PageEnding
	HTML PageEnding = "html"
	// CSS PageEnding
	CSS PageEnding = "css"
	// JPG PageEnding
	JPG PageEnding = "jpg"
	// PDF PageEnding
	PDF PageEnding = "pdf"
)

// MatchPageEnding from a string to a PageEnding
func MatchPageEnding(ending string) (PageEnding, error) {
	switch ending {
	case string(HTML):
		return HTML, nil
	case string(CSS):
		return CSS, nil
	case string(JPG):
		return JPG, nil
	case string(PDF):
		return PDF, nil
	default:
		return HTML, fmt.Errorf("unable to match PageEnding")
	}
}
