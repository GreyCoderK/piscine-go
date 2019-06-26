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
		listItem := l.Tail
        
        	//parcours pour se mettre a la fin de la list
        	for listItem.Next != nil {
            		listItem = listItem.Next
        	}
        
        	//enregistrement du dernier element a la liste
        	listItem.Next = n
        	l.Tail = n
	}
}
