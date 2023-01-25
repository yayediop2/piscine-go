package piscine

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _, v := range tab {
		if f(v) {
			for _, value := range v {
				if value >= '0' && value <= '9' {
					x = x + 1
				}
			}
		}
	}
	return x
}
