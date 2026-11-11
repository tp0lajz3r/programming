package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) > 1 {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[0] = 1
			n -= 1
		}
		if flowerbed[len(flowerbed)-2] == 0 && flowerbed[len(flowerbed)-1] == 0 {
			flowerbed[len(flowerbed)-1] = 1
			n -= 1
		}
	} else {
		if flowerbed[0] == 0 {
			flowerbed[0] = 1
			n -= 1
		}
	}

	for i := 2; i < len(flowerbed); i++ {
		if flowerbed[i-2] == 0 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
			flowerbed[i-1] = 1
			n -= 1
		}
	}

	if n > 0 {
		return false
	} else {
		return true
	}
}

func main() {
	// Case 1
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1

	// Case 2
	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := 2

	fmt.Println(canPlaceFlowers(flowerbed, n))
}
