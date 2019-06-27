package piscine

func ListReverse(l *List) {
	current := l.Head
        var precedent =  &NodeL{}

        for current != nil { 
                next := current.Next
                current.Next = precedent
                precedent = current
                current = next
        }
        temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
