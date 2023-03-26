package piscine

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	tete := l.Head
	var stock *interface{}
	for tete != nil {
		if comp(tete.Data, ref) {
			stock = &tete.Data
		}
		tete = tete.Next
	}
	return stock
}

func CompStr(a, b interface{}) bool {
	return a == b
}
