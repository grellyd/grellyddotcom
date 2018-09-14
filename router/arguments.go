package router

// waiiitttt a second. Isn't it always the last section?
// not absoultely. We can also have root/blog/item


// well that is root/section/page

// so we can walk backwards, and either give nothing as the section and page, nothing as the page, or all three.
// so pass in pointers to strings

// Argument is the portion of the routing url which the argument belongs
type Argument int

const (
	// NOARGS used from the routing url.
	NOARGS = Argument(iota)
	// FIRST portion of the routing url. EG: root/arg
	FIRST 
	// SECOND portion of the routing url. EG: root/topic/arg/
	SECOND
)
