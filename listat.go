package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Next == nil && pos > 0 {
		return nil
	} else {
		for pos != 0 {
			if l.Next == nil {
				return nil
			}
			l = l.Next
			pos--
		}
	}
	return l
}
