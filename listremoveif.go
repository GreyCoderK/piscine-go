package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}
	current := l.Head
	var previous *NodeL = nil
 
	for current != nil {
		if current.Data == data_ref {
			if previous == nil {
				l.Head = current.Next 	
			}else{
				previous = current
				previous.Next = current.Next
			}
		}else{
			previous = current
		}
		current = current.Next			
	}
}
