package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	tableau := make([]int, max)
	for i := range tableau {
		tableau[i] = i
	}
	return tableau[min:max]
}
