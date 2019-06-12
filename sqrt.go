package piscine

func Sqrt(nb int) {
	if nb < 0 {
		return 0
	}
	for i := 0; i < 101; i++ {
		if nb == i*i {
			nb = i
			break
		} else if i*i > nb {
			nb = 0
			break
		}
	}
	return nb
}
