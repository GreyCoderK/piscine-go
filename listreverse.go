package piscine

func ListReverse(l *List) {
	current := l.Head
	var prev *NodeL

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
