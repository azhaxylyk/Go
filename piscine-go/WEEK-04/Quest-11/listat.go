package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	iterator := l
	inc := 0
	for iterator != nil {
		if pos == inc {
			return iterator
		}
		inc++
		iterator = iterator.Next
	}
	return nil
}
