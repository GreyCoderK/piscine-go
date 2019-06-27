package piscine

func ListAt(l *NodeL, pos int) *NodeL{
	if l == nil {
		return nil 
	}
	if pos == 1 {
		return l
	}else{
		cmpt := 1
		current := l
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
