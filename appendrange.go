package piscine


func AppendRange(min, max int) []int {
	if min >= max {
		return []int
	}else{
		tab := [max-min]int
		for i:= min;min <max;i++ {
			tab = append(i)
		}
		return tab
	}
}

