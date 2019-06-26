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
		l.Tail=n
	}else{
		tail := l.Tail
		tail.Next=n
	}
}
