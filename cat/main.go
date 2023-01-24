package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	arg := os.Args
	if len(arg) > 1 {
		for i := 1; i < len(arg); i++ {
			content, err := ioutil.ReadFile(arg[i])
			if err != nil {
				PrintStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			PrintStr(string(content))
		}
	} else if len(arg) == 1 {
		for i := 1; i < len(arg); i++ {
			content, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				PrintStr("ERROR: " + err.Error())
				os.Exit(1)
			}
			PrintStr(string(content))
		}
	}
}
