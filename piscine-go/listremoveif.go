package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	pointer := l.Head
	for pointer != nil {
		if pointer.Data == data_ref {
			if l.Head == pointer {
				l.Head = pointer.Next
			} else {
				prev.Next = pointer.Next
			}
		}
		if pointer.Data != data_ref {
			prev = pointer
		}
		pointer = pointer.Next
	}
}
