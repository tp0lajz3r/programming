package main

import "fmt"

func productExceptSelf(nums []int) []int {
	leftP := make([]int, len(nums))
	rightP := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		leftP[i] = 1
		rightP[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		leftP[i] = nums[i-1] * leftP[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightP[i] = nums[i+1] * rightP[i+1]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = leftP[i] * rightP[i]
	}
	return nums
}

func main() {
	// Case 1
	nums := []int{1, 2, 3, 4}

	// Case 2
	// nums := []int{-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf(nums))
}
