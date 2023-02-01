package piscine

func ListClear(l *List) {
	if l.Head != nil {
		l.Head.Next = nil
		l.Head = nil
	}
}
