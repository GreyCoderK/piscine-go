package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
        cmpt := 0
	next := l.Next
	var first,temp *NodeI

        for current != nil {
		if cmpt == 0 {
			first = current
			cmpt++
			temp = next.Next
		}
		if cmpt > 0 {
			if temp == nil {
				return first
			}else{
				next = temp
			}
		}
		for next != nil {
                        if current.Data > next.Data {
                                current.Data, next.Data = next.Data,current.Data	
			}
                        next = next.Next
                }
                current = current.Next
		temp = current.Next
        }
	return first
}
