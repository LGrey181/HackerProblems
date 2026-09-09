package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}
	}
	fmt.Println(ret)

}

func rotate(s rune, delta int, key []rune) rune {
	ind := strings.IndexRune(string(key), s)
	for i, r := range key {
		if r == s {
			ind = i
			break
		}
	}
	if ind < 0 {
		panic("ind < 0")
	}
	ind = (ind + delta) % len(key)
	return key[ind]
}
