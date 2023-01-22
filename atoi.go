package piscine

func Atoi(s string) int {
	sr := []rune(s)
	var x int = 0

	if sr[0] == '+' {
		for i := 1; i < len(sr); i++ {
			if sr[i] < '0' || sr[i] > '9' {
				return 0
			}
			x = x*10 + int(sr[i]-'0')
		}
	} else if sr[0] == '-' {
		for i := 1; i < len(sr); i++ {
			if sr[i] < '0' || sr[i] > '9' {
				return 0
			}
			x = x*10 + int(sr[i]-'0')
		}
		x = -1 * x
	} else {
		for i := 0; i < len(sr); i++ {
			if sr[i] < '0' || sr[i] > '9' {
				return 0
			}
			x = x*10 + int(sr[i]-'0')
		}
	}
	return x
}
