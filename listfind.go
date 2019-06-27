package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head 
	for current != nil {
		if CompStr(ref, current.Data) {
			return &ref
		}
		current = current.Next
	}
	return nil
}
