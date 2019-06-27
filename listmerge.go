package piscine

func ListMerge(l1, l2 *List) {
	current := l1.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = l2.Head
}
