package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI{
	n := &NodeI{Data:data_ref}
        if l == nil {                   
                return n
        }
        first := l
        for l.Next != nil { 
                l = l.Next
        }
        l.Next = n

        return ListSort(first)
}
