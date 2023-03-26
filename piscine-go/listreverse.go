package piscine

func ListReverse(l *List) {
	var prev *NodeL
	var next *NodeL
	for l.Head != nil {
		next = l.Head.Next
		l.Head.Next = prev
		prev = l.Head
		l.Head = next
	}
	l.Head = prev
}
