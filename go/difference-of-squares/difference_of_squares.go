package diffsquares

// Difference the difference between the square of the sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

/* SquareOfSum the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
*/

// SquareOfSum without loop
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

/* SumOfSquares the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}
*/

// SumOfSquares using recursive
func SumOfSquares(n int) int {
	if n == 0 {
		return 0
	}

	return n*n + SumOfSquares(n-1)
}
