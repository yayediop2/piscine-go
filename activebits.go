package piscine

func ActiveBits(n int) int {
	c := 0
	if n < 1 {
		return c
	}
	mod := n % 2
	div := n / 2
	return mod + ActiveBits(div)
}
