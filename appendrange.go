package piscine

func AppendRange(min, max int) []int {
	if int(min) >= int(max) {
		var tab []int = nil
		return tab
	}else{
		tab := []int{}
		for i:= int(min);int(min) <int(max);i++ {
			tab = append(tab,i)
		}
		return tab
	}
}

