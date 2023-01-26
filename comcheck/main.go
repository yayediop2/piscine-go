package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, v := range arg {
		if v == "galaxy" || v == "01" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
