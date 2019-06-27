package piscine

func ListSize(l *List) int {
	if l.Head == nil{
		return 0
	}else{
		current := l.Head
		cmpt := 1
		for current.Next != nil {
			current = current.Next
			cmpt++
		}
		return cmpt
	}
}
