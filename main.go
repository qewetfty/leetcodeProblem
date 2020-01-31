package main

import "fmt"

func moveZeroes(nums []int) {
	lastZeroPosition := 0
	for i, num := range nums {
		if num != 0 {
			nums[lastZeroPosition] = nums[i]
			lastZeroPosition++
		}
	}
	for lastZeroPosition < len(nums) {
		nums[lastZeroPosition] = 0
		lastZeroPosition++
	}
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
