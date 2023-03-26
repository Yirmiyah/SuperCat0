package piscine

func Join(strs []string, sep string) string {
	var array string
	for i := 0; i < len(strs); i++ {
		if strs[i] != strs[len(strs)-1] {
			strs[i] += sep
		}
		array += strs[i]
	}
	return array
}
