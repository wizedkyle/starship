package utils

import "github.com/microcosm-cc/bluemonday"

var HtmlSanitizer *bluemonday.Policy

// SanitizerInit
// Creates an instance of bluemonday for sanitizing user input.
func SanitizerInit() {
	HtmlSanitizer = bluemonday.UGCPolicy()
}
