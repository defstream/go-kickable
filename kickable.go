// Package Kickable is built to answer the age old question, Can I Kick It?
// Kickable provides a wide variety of implementations as well.
//
// See README.md for more info.
package kickable

import (
	"strings"
)

const (
	// No is the message returned when an item is not kickable.
	No = "No."
	// Yes is the message returned when an item is kickable.
	Yes = "Yes, yes you can."
)

// Kickable returns true if the passed in string is considered kickable
func Kickable(s string) bool {
	if strings.EqualFold(s, "it") {
		return true
	}
	return false
}

// CanIKick returns a string response to the question, Can I Kick "it".
func CanIKick(it string) string {
	if Kickable(it) {
		return Yes
	}
	return No
}
