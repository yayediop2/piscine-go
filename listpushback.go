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
	n := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
