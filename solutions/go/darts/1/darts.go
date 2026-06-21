package darts

func Score(x, y float64) int {
	radSqr := x*x + y*y
    if radSqr <= 1 {
        return 10
    }
	if radSqr <= 25 {
        return 5
    }
	if radSqr <= 100 {
        return 1
    }
	return 0
}
