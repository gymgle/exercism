package diffsquares

// Difference the difference between the square of the sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareOfSum the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}
