package array

import "fmt"

// move array costs times
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			nums = append(nums[:j], nums[j+1:]...)
		} else {
			i++
			j++
		}
	}
	return len(nums)
}

// two pointer method, faster than first method
func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	first := 0
	second := 0
	prev := nums[0]
	for first < len(nums) {
		currNum := nums[first]
		if prev != currNum {
			prev = currNum
			second++
			nums[second] = currNum
		}
		first++
	}
	return second + 1
}

func testProblem26() {
	input1 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(input1))
	input2 := []int{0, 0, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(input2))
}
