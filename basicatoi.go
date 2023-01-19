package piscine

func BasicAtoi(s string) int {
	str := []rune(s)
	x := 0
	for i := 0; i < len(str); i++ {
		x = x*10 + int(str[i]-'0')
	}
	return x
}
