package piscine

func RecursiveFactorial(nb int) int {
	var r int = 1
	if nb > 0 {
		r = nb * RecursiveFactorial(nb-1)
	}

	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 20 {
		return 0
	}
	return r
}
