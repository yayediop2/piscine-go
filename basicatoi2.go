package main

import "fmt"

func BasicAtoi2(s string) int {
	sr := []rune(s)
	var x int = 0
	for _, value := range sr {
		if value >= '0' && value <= '9' {
			x = x*10 + int(value-'0')
		} else {
			return 0
		}
	}
	return x
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
