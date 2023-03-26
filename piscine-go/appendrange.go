package piscine

func AppendRange(min int, max int) []int {
	char := []int{}
	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		char = append(char, i)
	}
	return char
}
