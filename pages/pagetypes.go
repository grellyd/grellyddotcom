package pages

// PageType is the type of a page indicating how the page should be treated
type PageType string 

const (
	// STATIC PageType
	STATIC PageType = "static"
	// BLOG PageType
	BLOG PageType = "blog"
)
