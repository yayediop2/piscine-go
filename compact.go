package piscine

const N = 6

func Compact(ptr *[]string) int {
	var t []string
	for _, v := range *ptr {
		if v != "" {
			t = append(t, v)
		}
		*ptr = t
	}
	return len(t)
}
