package array

import (
	"fmt"
	"math"
)

// Given a binary array, find the maximum number of consecutive 1s in this array.

func test() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
	nums2 := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(nums2))
	nums3 := []int{0, 0}
	fmt.Println(findMaxConsecutiveOnes(nums3))
	nums4 := []int{1}
	fmt.Println(findMaxConsecutiveOnes(nums4))
}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	oneNum := 0
	for _, num := range nums {
		if num == 1 {
			oneNum++
		} else {
			max = int(math.Max(float64(oneNum), float64(max)))
			oneNum = 0
		}
	}
	max = int(math.Max(float64(oneNum), float64(max)))
	return max
}
