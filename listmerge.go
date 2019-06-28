package piscine

func ListMerge(l1, l2 *List) {
	
	if l1 == nil || l2 == nil{
		return
	}

	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Head
		return
	}

	current := l1.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = l2.Head
}
