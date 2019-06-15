package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	}else{
		tab := [max-min]int{}
		for i:= min;min <max;i++ {
			tab = append(tab,i)
		}
		return tab
	}
}

