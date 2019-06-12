package piscine

func IterativeFactorial(x int) int {
	if 0 == x || 1 == x {
		return 1
	} else {
		res := 1
		for i := 2; i <= x; i++ {
			res *= i
		}
		return res
	}
}
