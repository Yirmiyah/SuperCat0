package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head, l1.Tail = l2.Head, l2.Tail
		return
	} else {
		l1.Tail.Next = l2.Head
	}
}
