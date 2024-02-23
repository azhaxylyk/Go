package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	l.Tail, l.Head = l.Head, l.Tail

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
}
