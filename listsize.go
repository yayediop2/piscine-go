package piscine

func ListSize(l *List) int {
	c := 0
	for l.Head != nil {
		c++
		l.Head = l.Head.Next
	}
	return c
}
