package piscine

func CollatzCountDown(start int) int {
	c := 0
	if start <= 0 {
		return -1
	}
	for start > 1 {
		if start%2 == 0 {
			start = (start / 2)
		} else if start%2 == 1 {
			start = ((start * 3) + 1)
		}
		c++
	}
	return c
}

/* func main() {
	steps := CollatzCountDown(12)
	fmt.Println(steps)
}
*/
