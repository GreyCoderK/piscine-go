package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}

	current := l
        cmpt := 0
	next := l.Next
	var temp *NodeI
	
	if next == nil{
		return  l 
	}

        for current != nil {
		if cmpt == 0 {
			cmpt++
		}
		if cmpt > 0 {
			if temp == nil {
				return current
			}else{
				current = current.Next
				next = current.Next
			}
		}
		for next != nil {
                        if current.Data > next.Data {
                                current.Data, next.Data = next.Data,current.Data	
			}
                        next = next.Next
                }
		temp = current.Next
        }
	return current
}
