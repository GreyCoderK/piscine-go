package piscine

func CollatzCountdown(start int) int {
	step := 1
	for start > 1 {
		if start % 2 == 1 {
			start = (start * 3) + 1
		}else{
			start /= 2
		}
		step++
	}
	return step
}
