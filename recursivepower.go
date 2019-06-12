package piscine

func RecursivePower(nb, power int) int {
	if 1 == power {
		return nb
	} else if 0 == power {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
