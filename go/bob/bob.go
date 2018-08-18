// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var response string
	remark = strings.TrimSpace(remark)
	switch {
	case remark == "":
		response = "Fine. Be that way!"
	case remark == strings.ToUpper(remark) && remark != strings.ToLower(remark) && strings.HasSuffix(remark, "?"):
		response = "Calm down, I know what I'm doing!"
	case remark == strings.ToUpper(remark) && remark != strings.ToLower(remark):
		response = "Whoa, chill out!"
	case strings.HasSuffix(remark, "?"):
		response = "Sure."
	default:
		response = "Whatever."
	}
	return response
}
