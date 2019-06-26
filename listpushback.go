package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n:=&NodeL{Data:data}
        if l.Head==nil{
                l.Head=n
        }else if l.Tail==nil{
                head := l.Head
		head.Next = n
		l.Tail = n 
        }else{
		l.Tail.Next = n
	}
}
