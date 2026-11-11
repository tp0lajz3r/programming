package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	t := strings.Fields(s)
	slices.Reverse(t)
	return strings.Join(t, " ")
}

func main() {
	// case 1
	s := "the sky is blue"

	// case 2
	// s := "  hello world  "

	// case 3
	// s := "a good   example"

	fmt.Println(reverseWords(s))
}
