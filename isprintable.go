package piscine

func IsPrintable(s string) bool {
	bs := []byte(s)
	for _, value := range bs {
		if value <= 31 {
			return false
		}
	}
	return true
}
