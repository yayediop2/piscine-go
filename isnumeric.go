package main

import (
	"fmt"
)

func IsNumeric(s string) bool {
	bs := []byte(s)
	for _, value := range bs {
		if value < 48 || value > 57 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
