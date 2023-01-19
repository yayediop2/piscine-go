package main

import "fmt"

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	z := max - min
	a := make([]int, z)

	for i := 0; i < z; i++ {
		a[i] = min + i
	}
	return a
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
