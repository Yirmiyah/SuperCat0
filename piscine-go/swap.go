package piscine

func Swap(a *int, b *int) {
	change := *a
	*a = *b
	*b = change
}
