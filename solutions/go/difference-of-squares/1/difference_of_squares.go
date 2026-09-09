package differenceofsquares

func SquareOfSum(n int) int {
	summ := (n * (n + 1)) / 2
	return summ * summ
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
