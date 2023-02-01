package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Next == nil && pos > 0 {
		return nil
	}
	for l != nil {
		l = l.Next
		pos--
	}
	return l
}
