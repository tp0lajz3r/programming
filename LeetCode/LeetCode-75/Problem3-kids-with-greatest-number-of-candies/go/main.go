package main

import (
	"fmt"
	"sort"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	sorted_candies := make([]int, len(candies))
	copy(sorted_candies, candies)
	sort.Ints(sorted_candies)
	max := sorted_candies[len(sorted_candies)-1] - extraCandies
	res := []bool{}
	for i := 0; i < len(candies); i++ {
		if candies[i] < max {
			res = append(res, false)
		} else {
			res = append(res, true)
		}
	}
	return res

}

func main() {
	// case 1
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	// case 2
	// candies := []int{4, 2, 1, 1, 2}
	// extraCandies = 1

	// case 3
	// candies := []int{12, 1, 12}
	// extraCandies = 10

	fmt.Println((kidsWithCandies(candies, extraCandies)))
}
