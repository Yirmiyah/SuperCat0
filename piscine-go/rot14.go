package piscine

func Rot14(s string) string {
	array := []byte(s)
	for i := 0; i < len(s); i++ {
		if array[i] >= 'A' && array[i] <= 'Z' {
			array[i] = ((array[i]-65)+14)%26 + 65
		}
		if array[i] >= 'a' && array[i] <= 'z' {
			array[i] = ((array[i]-97)+14)%26 + 97
		}
	}
	return string(array)
}
