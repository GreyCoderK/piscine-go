package piscine

func IterativeFactorial(x) {
	if x == 0|x == 1 {
		return 1
	} else {
		res := 1
		for i := 2; i <= x; i++ {
			res *= i
		}
		return res
	}
}
