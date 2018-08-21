package hamming

import (
	"errors"
)

// Distance calculate the Hamming difference between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("different lengths between two sequences")
	}

	var hamming int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming++
		}
	}
	return hamming, nil
}
