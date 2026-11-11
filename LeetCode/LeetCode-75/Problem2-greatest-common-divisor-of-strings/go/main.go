package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	a, b := len(str1), len(str2)
	if (str1 + str2) != (str2 + str1) {
		return ""
	} else {
		for b != 0 {
			a, b = b, a%b
		}
		return str1[:a]
	}
}

func main() {
	// Case 1
	str1 := "ABCABC"
	str2 := "ABC"

	// Case 2
	// str1 := "ABABAB"
	// str2 := "ABAB"

	// Case 3
	// str1 := "LEET"
	// str2 := "CODE"

	// Case 4
	// str1 := "AAAAAB"
	// str2 := "AAA"

	fmt.Println(gcdOfStrings(str1, str2))
}
