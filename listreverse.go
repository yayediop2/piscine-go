package piscine

func ListReverse(l *List) {
	var nod *NodeL = nil // c'est à dire qu'elle ne pointe vers auncun noeud
	for l.Head != nil {
		suiv := l.Head.Next
		l.Head.Next = nod
		nod = l.Head
		l.Head = suiv
	}
	l.Tail = l.Head
	l.Head = nod
}
