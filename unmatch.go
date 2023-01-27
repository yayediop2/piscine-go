package piscine

func Unmatch(a []int) int {
	a = TriInt(a)
	for i := 0; i < len(a); {
		if i < len(a)-1 && a[i] == a[i+1] {
			i += 2
		} else {
			return a[i]
		}
	}
	return -1
}

func TriInt(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
