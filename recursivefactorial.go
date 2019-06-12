package piscine

import "math"

func RecursiveFactorial(x int) int {
	if 0 == x || 1 == x {
		return 1
	} else if x*RecursiveFactorial(x-1) > math.MaxInt32 {
		return 0
	} else {
		return x * RecursiveFactorial(x-1)
	}
}
