package main

import "fmt"

func Atoi(s string) int {
	sr := []rune(s)
	var x int = 0

	if sr[0] == '+' {
		for i := 1; i < len(sr); i++ {
			if sr[i] < '0' || sr[i] > '9' {
				return 0
			}
			x = x*10 + int(sr[i]-'0')
		}
	} else if sr[0] == '-' {
		for i := 1; i < len(sr); i++ {
			if sr[i] < '0' || sr[i] > '9' {
				return 0
			}
			x = x*10 + int(sr[i]-'0')
		}
		x = -1 * x
	} else {
		for i := 0; i < len(sr); i++ {
			if sr[i] < '0' || sr[i] > '9' {
				return 0
			}
			x = x*10 + int(sr[i]-'0')
		}
	}
	return x
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
