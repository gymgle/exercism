package grains

import (
	"errors"
)

// Square returns the number of grains on a square on a chess board where the
// first square has 1 and every subsequent square doubles the number
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("illegal input")
	}

	return 1 << uint(n-1), nil
}

// Total returns the total number of grains on the chess board
func Total() uint64 {
	return (1 << 64) - 1
}
