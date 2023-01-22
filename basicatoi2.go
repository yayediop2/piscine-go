package piscine

func BasicAtoi2(s string) int {
	sr := []rune(s)
	var x int = 0
	for _, value := range sr {
		if value >= '0' && value <= '9' {
			x = x*10 + int(value-'0')
		} else {
			return 0
		}
	}
	return x
}
