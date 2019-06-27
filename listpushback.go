package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n:=&NodeL{Data:data}
        if l.Head==nil{
                l.Head=n
        }else{
		listItem := l.Head
        
        	//parcours pour se mettre a la fin de la list
        	for listItem.Next != nil {
            		listItem = listItem.Next
        	}
        
        	//enregistrement du dernier element a la liste
        	listItem.Next = n
	}
}
