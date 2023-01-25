package piscine

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _, v := range tab {
		if f(v) {
			x = x + 1
		}
	}
	return x
}
