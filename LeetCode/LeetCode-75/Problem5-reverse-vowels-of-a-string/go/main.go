package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseVowels(s string) string {
	vowels := [5]string{"a", "e", "i", "o", "u"}
	tmp := []string{}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(vowels); j++ {
			if vowels[j] == strings.ToLower(string(s[i])) {
				tmp = append(tmp, string(s[i]))
			}
		}
	}
	slices.Reverse(tmp)
	l := strings.Split(s, "")
	c := 0
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(vowels); j++ {
			if vowels[j] == strings.ToLower(string(l[i])) {
				l[i] = tmp[c]
				c++
				break
			}
		}
	}
	return strings.Join(l, "")
}

func main() {
	// Case 1
	s := "IceCreAm"

	fmt.Println(reverseVowels(s))
}
