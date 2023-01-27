package piscine

func Join(strs []string, sep string) string {
	var s string
	for i := 0; i < len(strs); i++ {
		s = s + strs[i]
		if i < len(strs)-1 {
			s += sep
		}
	}
	return s
}
