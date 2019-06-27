package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}else{
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		return current.Data
	}
}
