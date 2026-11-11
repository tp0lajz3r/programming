package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	if len(word1) >= len(word2) {
		for i := 0; i < len(word2); i++ {
			res += string(word1[i]) + string(word2[i])
		}
		res += string(word1[len(word2):])
	} else {
		for i := 0; i < len(word1); i++ {
			res += string(word1[i]) + string(word2[i])
		}
		res += string(word2[len(word1):])
	}
	return res
}

func main() {
	// Case 1
	word1 := "abc"
	word2 := "pqr"

	// Case 2
	// word1 := "ab"
	// word2 := "pqrs"

	// Case 3
	// word1 := "abcd"
	// word2 := "pq"

	fmt.Println(mergeAlternately(word1, word2))
}
