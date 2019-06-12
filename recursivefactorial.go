package piscine

func RecursiveFactorial(x int) int {
	if 0 == x || 1 == x {
		return 1
	} else {
		return x * RecursiveFactorial(x-1)
	}
}
