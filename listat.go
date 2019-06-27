package piscine

func ListAt(l *NodeL, pos int) *NodeL{
	if l == nil {
		return nil
	}
	current := l
	index := 0
	for index < pos {
		if current.Next == nil {
			return nil
		}
		current = current.Next
		index++
	}
	return current
}
