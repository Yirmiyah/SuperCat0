package piscine

func Unmatch(a []int) int {
	var paire int
	paire = 0
	for _, oneloop := range a {
		for _, twoloop := range a {
			if twoloop == oneloop {
				paire++
			}
		}
		if paire%2 != 0 {
			return oneloop
		}
	}
	return -1
}
