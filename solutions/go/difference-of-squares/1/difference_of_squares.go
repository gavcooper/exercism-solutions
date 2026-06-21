package diffsquares

func SquareOfSum(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
    for i := 1; i <= n; i++ {
        sum += i * i
    }
	return sum
}

func Difference(n int) int {
    if SumOfSquares(n) >= SquareOfSum(n) {
        return SumOfSquares(n) - SquareOfSum(n)
    }
	return SquareOfSum(n) - SumOfSquares(n)
}
