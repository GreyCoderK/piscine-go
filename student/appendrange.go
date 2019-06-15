package piscine


func AppendRange(min, max int) []int {
	if min >= max {
		return []
	}else{
		tab := make([]int, max-min)
		for i:= min;min <max+1;i++ {
			tab = append(i)
		}
		return tab
	}
}

