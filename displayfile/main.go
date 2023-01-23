package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) == 2 {
		fichier, err := ioutil.ReadFile(arg[1])
		if err != nil {
			fmt.Printf("ERREUR!%v\n", err.Error())
		}
		fmt.Print(string(fichier))
	} else if len(arg) == 1 {
		println("File name missing")
	} else if len(arg) > 2 {
		fmt.Println("Too many arguments")
	}
}
