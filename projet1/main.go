package main

import (
	"fmt"
	"strconv"
)

func main() {
	//hex
func hexadecimal(s string) int{
	var n=0
	for i:=0; i<len(s); i++{
		n*=16
		if s[i]>='0' && s[i]<='9'{
			n+=int(s[i]-'0')
		} else if s[i]>='a' && s[i]<='z'{
			n+=int(s[i]-'a'+10)
		}else if s[i]>='A' && s[i]<='Z'{
			n+=int(s[i]-'A'+10)
		}
	}
	return n
}

func 
	/*a, _ := strconv.ParseInt("1E", 16, 32)
	//bin
	b, _ := strconv.ParseInt("11", 2, 32) */
	fmt.Println(b)
	fmt.Println(a)
}
