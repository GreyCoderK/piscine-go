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
	//var first,temp *NodeI
	
	if next == nil{
		return  l 
	}

        for current != nil {
		//if cmpt == 0 {
			//first = current
		//	cmpt++
		//}
		//if cmpt > 0 {
			//if temp == nil {
			//	return first
			//}else{
			//	next = temp
			//}
		//}
		for next != nil {
                        if current.Data > next.Data {
                                current.Data, next.Data = next.Data,current.Data	
			}
                        next = next.Next
                }
                current = current.Next
		//temp = current.Next
        }
	return current
}
