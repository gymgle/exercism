package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determine if a word or phrase is an isogram
func IsIsogram(s string) bool {
	exist := make(map[rune]bool)
	s = strings.ToUpper(s)
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			if _, ok := exist[letter]; !ok {
				exist[letter] = true
			} else {
				return false
			}
		}
	}
	return true
}
