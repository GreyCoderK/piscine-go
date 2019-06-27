package piscine

func ListAt(l *NodeL, pos int) *NodeL{
	if pos == 1 {
		return l.Head
	}else{
		cmpt := 1
		current := l.Head
		for current.Next != nil {
			current = current.Next
			cmpt++
			if cmpt == pos {
				return current
			}
		}
	}
	return nil
}
