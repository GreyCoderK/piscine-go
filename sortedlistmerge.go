package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return ListSort(n2)
	}
	if n2 == nil { 
                return ListSort(n1)
        }
	first := n1
	for n1.Next != nil {
		n1 = n1.Next
	}
	n1.Next = n2
	return ListSort(first)
}
