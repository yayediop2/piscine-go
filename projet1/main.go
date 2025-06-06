package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func hexadecimal(s string) string {
	var n = 0
	for i := 0; i < len(s); i++ {
		n *= 16
		if s[i] >= '0' && s[i] <= '9' {
			n += int(s[i] - '0')
		} else if s[i] >= 'a' && s[i] <= 'z' {
			n += int(s[i] - 'a' + 10)
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			n += int(s[i] - 'A' + 10)
		}
	}
	str := strconv.Itoa(n)
	return string(str)
}

func binaire(s string) string {
	n := 0
	for _, l := range s {
		n *= 2
		if l == '0' || l == '1' {
			n += int(l - '0')
		}
	}
	str := strconv.Itoa(n)
	return string(str)
}

func ToUpper(s string) string {
	bs := []rune(s)
	for i, value := range bs {
		if value >= 97 && value <= 122 {
			bs[i] -= 32
		}
	}
	return string(bs)
}

func ToLower(s string) string {
	bs := []rune(s)
	for i, value := range bs {
		if value >= 65 && value <= 90 {
			bs[i] += 32
		}
	}
	return string(bs)
}

func Capitalize(s string) string {
	sr := []rune(s)
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if i == 0 && s[i] >= 'a' && s[i] <= 'z' {
				sr[i] = ((sr[i]) - ' ')
				/* if sr[i] == ' ' {
					sr[i+1] = (sr[i+1] + ' ')
				} */
			}
		}
	}
	return string(sr)
}

func SplitWhiteSpaces(s string) []string {
	tab := []string{}
	m := ""
	for _, letter := range s {
		if letter != ' ' && letter > 33 {
			m += string(letter)
		}
		if letter == ' ' || letter == rune(s[len(s)-1]) {
			if m != "" {
				tab = append(tab, m)
				m = "" // ENLEVER çA APRES -> ""
			}
		}
	}
	return tab
}

func recupF() []string {
	arg := os.Args[1:]
	var o []string
	for i := 0; i < len(arg); i++ {
		fich1, _ := ioutil.ReadFile(arg[0])
		fich2 := ioutil.WriteFile(arg[1], []byte(Join(o)), 0777)
		if fich2 != nil {
			panic(fich2)
		}
		str := string(fich1)
		c := SplitWhiteSpaces(string(str))
		o = motcle(c)
		//fmt.Print(o)
		//fmt.Print(c)
	}
	return o
}

func Join(strs []string) string { // utiliser join pttr
	var s string
	for i := 0; i < len(strs); i++ {
		s = s + strs[i]
		if i < len(strs)-1 {
			s += " "
		}
	}
	return s
}

func motcle(s []string) []string {
	tableau := []string{}
	for i := 0; i < len(s); i++ {
		var num int

		//hex
		if s[i] == "(hex)" {
			s[i-1] = hexadecimal(s[i-1])
			s[i] = ""
		}
		//bin
		if s[i] == "(bin)" {
			s[i-1] = binaire(s[i-1])
			s[i] = ""
		}
		//cap
		if s[i] == "(cap)" {
			s[i-1] = Capitalize(s[i-1])
			s[i] = ""
		}
		// (cap,
		if s[i] == "(cap," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			s[i], s[i+1] = "", ""
			for p := 1; p <= num; p++ {
				s[i-num] = Capitalize(s[i-num])
				i++
			}
		}
		//low
		if s[i] == "(low)" {
			s[i-1] = ToLower(s[i-1])
			s[i] = ""
		}
		// (low,
		if s[i] == "(low," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' { // gerer si on 14 par exemple. Eviter le out of range
					num = int(l - '0')
					break // on aurait plus besoin de ça
				}
			}
			s[i], s[i+1] = "", ""
			for p := 1; p <= num; p++ {
				s[i-num] = ToLower(s[i-num])
				i++
			}
		}
		//up
		if s[i] == "(up)" {
			s[i-1] = ToUpper(s[i-1])
			s[i] = ""
		}
		// (up,
		if s[i] == "(up," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			s[i], s[i+1] = "", ""
			for p := 1; p <= num; p++ {
				s[i-num] = ToUpper(s[i-num])
				i++
			}
		}
		if s[i] != "" {
			tableau = append(tableau, s[i]) // ICI LE TABLEAU N'EST LIÉ QU'A ATOAN
		}

	}
	//Ponctuation(s)
	AToAn(tableau)
	return tableau
}

/*
	 func Ponctuation(ss []string) []string {
		var st []string
		s := Join(ss)
		for i := 0; i < len(s); i++ {
		}
		return st
	}
*/
func AToAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "a" {
			if isVowel(string(s[i+1][0])) {
				s[i] = "an"
			}
		}
	}
	return s
}

func isVowel(s string) bool {
	v := []string{"a", "e", "u", "i", "o", "A", "E", "U", "I", "O"}
	for _, l := range v {
		if s == l {
			return true
		}
	}
	return false
}

func main() {
	recupF()
}
