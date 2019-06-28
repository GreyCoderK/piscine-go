package piscine

func ListMerge(l1, l2 *List) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	current := l1.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = l2.Head
}
