package piscine

func IterativeFactorial(x int) {
	if x == int(0)|x == int(1) {
		return 1
	} else {
		res := 1
		for i := 2; i <= x; i++ {
			res *= i
		}
		return res
	}
}
