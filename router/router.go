package router

import (
	"regexp"
)


// Register takes a url path and a valid regexp for the route. Returns an error if there is a problem
func Register(path string, validPaths string) error {
	r, err := buildRegex(validPaths)
	return err
}

func buildRegex(validPaths string) (r *regexp.Regexp, err error) {
	r, err = regexp.Compile(validPaths)
	return r, err
}

