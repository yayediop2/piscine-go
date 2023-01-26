package piscine

func Rot14(s string) string {
	var val rune
	rs := []rune(s)
	rus := []rune{}
	for _, v := range rs {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			if v >= 'a' && v <= 'l' || v >= 'A' && v <= 'L' {
				val = v + 14
				rus = append(rus, val)
			} else if v >= 'm' && v <= 'z' || v >= 'M' && v <= 'Z' {
				reste := 'z' - rune(v)
				val = 'a' + 13 - reste
				rus = append(rus, val)
			}
		} else if v == ' ' || v == '!' || v == '?' {
			rus = append(rus, v)
		}
	}
	return string(rus)
}

/*
func Rot14(s string) string {
	var val rune
	rs := []rune{}

	for _, v := range rs {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			if v >= 'a' && v <= 'l' || v >= 'A' && v <= 'L' {
				val = v + 14
			} else if v >= 'm' && v <= 'z' || v >= 'M' && v <= 'Z' {
				reste := 'z' - rune(v)
				val = 'a' + 13 - reste
			}
		} else if v == ' ' || v == '!' || v == '?' {
			val = v
		}
		rs = append(rs, val)
	}
	return string(rs)
}
*/
