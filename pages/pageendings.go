package pages

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
