package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/* func ToUpper(s string) string {
	runeS := []rune(s)
	for i, value := range runeS {
		if value >= 97 && value <= 122 {
			runeS[i] -= 32
		}
	}
	return string(runeS)
} */

func Capitalize(s string) string {
	runeS := []rune(s)
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if i == 0 && s[i] >= 'a' && s[i] <= 'z' {
				runeS[i] -= 32
			} else if i > 0 && s[i] >= 'A' && s[i] <= 'Z' {
				runeS[i] += 32
			}
		}
	}
	return string(runeS)
}

func motcle(string0 []string) []string {
	string1 := strings.Join(string0, " ")
	string2 := strings.ReplaceAll(string1, "(", " (")
	string3 := strings.ReplaceAll(string2, "( ", "(")
	string4 := strings.ReplaceAll(string3, ")", ") ")
	string5 := strings.ReplaceAll(string4, " )", ")")
	s := strings.Fields(string5)
	taille := len(s)

	for i := 0; i < taille; i++ {
		var num int
		//hex
		if s[i] == "(hex)" {
			if IsHexa(s[i-1]) {
				if taille == 0 {
					break
				}
				if i > 0 {
					a, _ := strconv.ParseInt(s[i-1], 16, 32)
					s[i-1] = strconv.Itoa(int(a))
				}
				s = append(s[:i], s[(i+1):]...)
				taille = taille - 1

				if i > 0 {
					i-- // si on l'enlève, ça ne marche plus
				}
			} else {
				fmt.Println("ERREUR: le nombre hexadécimal donné est INCORRECT")
				s = append(s[:i], s[(i+1):]...)
				taille = taille - 1

			}
		}
		//bin
		if s[i] == "(bin)" {
			if i > 0 {
				a, _ := strconv.ParseInt(s[i-1], 2, 32)
				s[i-1] = strconv.Itoa(int(a))
			}
			s = append(s[:i], s[(i+1):]...)
			taille = taille - 1

			if taille == 0 {
				break
			}
			if i > 0 {
				i--
			}
		}
		//cap
		if s[i] == "(cap)" {
			if i > 0 {
				s[i-1] = Capitalize(s[i-1])
			}
			s = append(s[:i], s[(i+1):]...)
			taille = taille - 1

			if taille == 0 {
				break
			}
			if i > 0 {
				i--
			}
		}
		//up
		if s[i] == "(up)" {
			if i > 0 {
				s[i-1] = strings.ToUpper(s[i-1])
			}
			s = append(s[:i], s[(i+1):]...)
			taille = taille - 1

			if taille == 0 {
				break
			}
			if i > 0 {
				i--
			}
		}
		//low
		if s[i] == "(low)" {
			if i > 0 {
				s[i-1] = strings.ToLower(s[i-1])
			}
			s = append(s[:i], s[(i+1):]...)
			taille--

			if taille == 0 {
				break
			}
			if i > 0 {
				i--
			}
		}
		//(cap,
		if s[i] == "(cap," {
			nextS := s[i+1]
			num, _ = strconv.Atoi(nextS[:len(nextS)-1])
			if i > 0 {
				if num < 0 {
					fmt.Println("ERREUR: vous avez donné un nombre négatif !!!")
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					continue
				}
				if num <= i {
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					for p := num - 1; p >= 0; p-- {
						if i >= 0 && i < taille {
							s[i-1] = Capitalize(s[i-1])
							i--
						}
					}
				} else if num > i && i < taille {
					num = i
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					for p := num - 1; p >= 0; p-- {
						s[i-1] = Capitalize(s[i-1])
						i--
					}
					continue
				}
			} else {
				s = append(s[:i], s[(i+2):]...)
				taille = taille - 2
			}

			if taille == 0 {
				break
			}
			if i > 0 {
				i--
			}
		}

		//(low,
		if s[i] == "(low," {
			nextS := s[i+1]
			num, _ = strconv.Atoi(nextS[:len(nextS)-1])
			if i > 0 {
				if num < 0 {
					fmt.Println("ERREUR: vous avez donné un nombre négatif !!!")
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					continue
				}
				if num <= i {
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					for p := num - 1; p >= 0; p-- {
						if i >= 0 && i < taille {
							s[i-1] = strings.ToLower(s[i-1])
							i--
						}
					}
				} else if num > i {
					num = i
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					for p := num - 1; p >= 0; p-- {
						s[i-1] = strings.ToLower(s[i-1])
						i--
					}
					continue
				}
			} else {
				s = append(s[:i], s[(i+2):]...)
				taille = taille - 2

			}
			if taille == 0 {
				break
			}
			if i > 0 {
				i--
			}
		}
		//(up,
		if s[i] == "(up," {
			nextS := s[i+1]
			num, _ = strconv.Atoi(nextS[:len(nextS)-1])
			if i > 0 {
				if num < 0 {
					fmt.Println("ERREUR: vous avez donné un nombre négatif !!!")
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					continue
				}
				if num <= i {
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					for p := num - 1; p >= 0; p-- {
						if i >= 0 && i < taille {
							s[i-1] = strings.ToUpper(s[i-1])
							i--
						}
					}
				} else if num > i {
					num = i
					s = append(s[:i], s[(i+2):]...)
					taille = taille - 2

					for p := num - 1; p >= 0; p-- {
						s[i-1] = strings.ToUpper(s[i-1])
						i--
					}
					continue
				}
			} else {
				s = append(s[:i], s[(i+2):]...)
				taille = taille - 2

			}
			if taille == 0 {
				break
			}
			if i > 0 {
				i--
			}
		}
	}
	return s
}

func IsHexa(s string) bool {
	h, err := regexp.MatchString("^[0-9a-fA-F]+$", s)
	return err == nil && h
}
func AToAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "a" && i < len(s)-1 {
			if isVowel(string(s[i+1][0])) {
				s[i] = "an"
			}
		}
		if s[i] == "A" && i < len(s)-1 {
			if isVowel(string(s[i+1][0])) {
				s[i] = "An"
			}
		}
	}
	return s
}
func isVowel(s string) bool {
	v := []string{"a", "e", "u", "i", "o", "h", "A", "E", "U", "I", "O", "H"}
	for _, v := range v {
		if s == v {
			return true
		}
	}
	return false
}

func apostorophiii(string0 []string) []string {
	string1 := strings.Join(string0, " ")
	string2 := strings.ReplaceAll(string1, "'", " ' ")
	string3 := strings.ReplaceAll(string2, "  ", " ")
	s := []rune(string3)
	A := 0
	taille := len(s)
	fmt.Println(string(s))
	if len(s) > 1 { // pour gérer le cas ou on a une seule apostrophe
		for i := 0; i < len(s); i++ {
			if (s[i] == '\'' || s[i] == '"') && (A == 0) {
				s = append(s[:i+1], s[(i+2):]...)
				taille--
				A = 1
				continue
			}
			if (s[i] == '\'' || s[i] == '"') && (A == 1) {
				s = append(s[:i-1], s[(i):]...)
				taille--
				A = 0
			}
			if s[len(s)-1] == '\'' || s[len(s)-1] == '"' {
				s = append(s[:i-1], s[(i):]...)
				taille--
			}
		}
	}
	slice := strings.Fields(string(s))
	return slice
}

func espacesSup(s string) string {
	for i, l := range s {
		if l == ' ' {
			i++
		}
	}
	return s
}

func EcrireF(s []string) {
	arg := os.Args[1:]
	fich2 := ioutil.WriteFile(arg[1], []byte(strings.Join(s, " ")), 0777)
	if fich2 != nil {
		panic(fich2)
	}
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Ce test est illogique")
		}
	}()
	arg := os.Args[1:]
	var a []string
	if len(arg) == 2 {
		for i := 0; i < len(arg); i++ {
			fich1, _ := ioutil.ReadFile(arg[0])
			a = strings.Fields(string(fich1))
		}
		fmt.Println(a)
		b := apostorophiii(a)
		c := AToAn(b)
		d := motcle(c)
		fmt.Println(d)
		EcrireF(d)
	}
}
