package piscine


func MakedRange(min, max int) []int {
        if min >= max {
                return []
        }else{
                tab := make([]int,max-min)
		j:=0
                for i:= min;min <max;i++ {
                        tab[j] = i
			j++
                }
                return tab
        }
}

