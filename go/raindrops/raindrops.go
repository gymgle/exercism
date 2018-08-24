package raindrops

import (
	"strconv"
)

// Convert a number to a string, the contents of which depend on the number's factors
func Convert(n int) string {
	var result string

	if n%3 == 0 {
		result += "Pling"
	}

	if n%5 == 0 {
		result += "Plang"
	}

	if n%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(n)
	}

	return result
}
