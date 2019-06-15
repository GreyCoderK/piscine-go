package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	}else{
		const size = max-min
		tab := [size]int
		for i:= min;min <max;i++ {
			tab = append(i)
		}
		return tab
	}
}

