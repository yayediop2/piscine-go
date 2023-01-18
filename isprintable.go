package main

import "fmt"

func IsPrintable(s string) bool {
	bs := []byte(s)
	for _, value := range bs {
		if value <= 31 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}
