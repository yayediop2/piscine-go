package piscine

func Capitalize(s string) string {
	sr := []rune(s)
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if sr[i] == ' ' {
				sr[i+1] = (sr[i+1] + ' ')
			}
		}
	}
	return string(sr)
}
