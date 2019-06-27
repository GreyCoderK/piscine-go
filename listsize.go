package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	if l.head == nil{
		return 0
	}else{
		current := l.head
		cmpt := 1
		for current.Next != nil {
			current = current.next
			cmpt++
		}
		return cmpt
	}
}
