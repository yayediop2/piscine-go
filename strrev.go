package piscine

func StrRev(s string) string {
	var t string
	for i := len(s) - 1; i >= 0; i-- {
		t += string(s[i])
	}
	return t
}
