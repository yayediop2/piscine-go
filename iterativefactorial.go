package piscine

func IterativeFactorial(nb int) int {
	var r int = 1
	for i := 1; i <= nb; i++ {
		if nb == 0 {
			return 1
		}
		r = r * i
	}
	if nb < 0 {
		return 0
	}

	return int(r)
}
