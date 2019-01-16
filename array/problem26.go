package array

import "fmt"

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

func main2() {
	input1 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(input1))
	input2 := []int{0, 0, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(input2))
}
