package picine

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Data.Next == nil && pos > 0 {
		return nil
	}
	for l.Data != nil {
		l.Data = l.Data.Next
		pos--
	}
	return l.Data
}
