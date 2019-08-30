package phonenumber

import (
	"fmt"
	"unicode"
)

// Number produce a pure number without country code
func Number(s string) (string, error) {
	if len(s) < 10 {
		return "", fmt.Errorf("not enough numbers %s", s)
	}

	n := make([]rune, 0, len(s))
	for _, c := range s {
		if c == '+' || c == '-' || c == '(' || c == ')' || c == '.' || c == ' ' {
			continue
		} else if unicode.IsDigit(c) {
			n = append(n, c)
		} else {
			return "", fmt.Errorf("%s contains invalid charactor %q", s, c)
		}
	}

	clean := string(n)
	lens := len(clean)
	switch {
	case lens < 10:
		return "", fmt.Errorf("not enough numbers %s", clean)
	case lens == 10 && clean[0] >= '2' && clean[3] >= '2':
		return clean, nil
	case lens == 11 && clean[0] == '1' && clean[1] >= '2' && clean[4] >= '2':
		return clean[1:], nil
	}

	return "", fmt.Errorf("%s contains too many numbers", s)
}

// Format produce a format as `(613) 995-0253`
func Format(s string) (string, error) {
	n, err := Number(s)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}

// AreaCode return the area code
func AreaCode(s string) (string, error) {
	n, err := Number(s)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", n[:3]), nil
}
