package piscine

func Compact(ptr *[]string) int {
	var array []string
	count := 0
	for _, element := range *ptr {
		if element != "" {
			array = append(array, element)
			count++
		}
	}
	*ptr = array
	return count
}
