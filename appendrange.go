package piscine


func AppendRange(min, max int) []int {
	if min >= max {
		return []
	}else{
		const size = max-min
		tab := [size]int
		for i:= min;min <max;i++ {
			tab = append(i)
		}
		return tab
	}
}

