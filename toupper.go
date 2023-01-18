package piscine

func ToUpper(s string) string {
	bs := []byte(s)
	for _, value := range bs {
		if value >= 97 && value <= 122 {
			value = value - 32
		}
	}
	return string(bs)
}
