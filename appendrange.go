package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	}else{
		var diff int = max - min
		tab := [diff]int{}
		for i:= min;min <max;i++ {
			tab = append(tab,i)
		}
		return tab
	}
}

