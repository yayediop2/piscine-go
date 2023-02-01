package piscine

type List struct {
	Head *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Head.Next == nil && pos > 0 {
		return nil
	}
	for l.Head != nil {
		l.Head = l.Head.Next
		pos--
	}
	return l.Head
}
