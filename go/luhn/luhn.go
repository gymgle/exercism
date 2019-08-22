package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Valid given a number determine whether or not it is valid per the Luhn formula
func Valid(number string) bool {
	num := strings.ReplaceAll(number, " ", "")

	if len(num) <= 1 {
		return false
	}

	r := "^[0-9]+$"
	if matched, _ := regexp.MatchString(r, num); !matched {
		return false
	}

	sum := 0
	for i := len(num) - 1; i >= 0; i -= 2 {
		n, _ := strconv.Atoi(string(num[i]))
		sum += n
	}

	for i := len(num) - 2; i >= 0; i -= 2 {
		n, _ := strconv.Atoi(string(num[i]))
		if n*2 >9 {
			sum += n*2-9
		} else {
			sum += n*2
		}
	}

	return sum%10 == 0
}
