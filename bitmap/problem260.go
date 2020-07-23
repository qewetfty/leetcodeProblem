package bitmap

import "fmt"

// Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear
// exactly twice. Find the two elements that appear only once.
//	Example:
//		Input:  [1,2,1,3,2,5]
//		Output: [3,5]
//	Note:
//		The order of the result is not important. So in the above example, [5, 3] is also correct.
//		Your algorithm should run in linear runtime complexity. Could you implement it using only
//		constant space complexity?

func singleNumber(nums []int) []int {
	bitmap := 0
	for _, num := range nums {
		bitmap = bitmap ^ num
	}
	diff, x := bitmap&(-bitmap), 0
	for _, num := range nums {
		if num&diff != 0 {
			x = x ^ num
		}
	}
	return []int{x, bitmap ^ x}
}

func testProblem260() {
	fmt.Println(singleNumber([]int{2, 1, 2, 3, 4, 1}))
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}
