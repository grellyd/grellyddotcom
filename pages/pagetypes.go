package pages

// PageType is the type of a page indicating how the page should be treated
// TODO: Refactor. Messy with page ending, and no struct/combined statics. 
type PageType string 

const (
	// STATIC PageType
	STATIC PageType = "static"
	// BLOG PageType
	BLOG PageType = "blog"
	// CSSRESOURCE PageType
	CSSRESOURCE PageType = "css"
	// IMAGESRESOURCE PageType
	IMAGESRESOURCE PageType = "images"
)
