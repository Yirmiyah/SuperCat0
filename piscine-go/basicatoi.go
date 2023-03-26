package piscine

func BasicAtoi(s string) int {
	s0 := s
	n := 0
	if len(s) > 0 {
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0
			}
		}
		if len(s) > 0 {
			for _, ch := range []byte(s) {
				ch -= '0'
				if ch > 9 {
					return 0
				}
				n = n*10 + int(ch)
			}
			if s0[0] == '-' {
				n = -n
			}
		}
	}
	return n
}
