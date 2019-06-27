package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI{
	n := &NodeI{Data:data_ref} 
	if l == nil {
		l = n
		return l
	}
	first := l
	for l.Next != nil {
		l = l.Next
	}
	l.Next = n

	return ListSort(first)
}
