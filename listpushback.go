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
	// initialisation du noeud
	noeud := &NodeL{Data:data}
	
	//verification que la liste n'est pas nulle
	if l.Head == nil {
		l.Head = noeud
	}else{
		//sauvegarde du premier element
		listItem := l.Head
		
		//parcours pour se mettre a la fin de la list
		for listItem.Head != nil {
			listItem = listItem.Next
		}
		
		//enregistrement du dernier element a la liste
		iterator.Next = noeud
		l.Tail = noeud
	}
}
