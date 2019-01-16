package array

import "fmt"

func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 1
	}
	position := 0
	for position < numsLen {
		number := nums[position]
		if number <= 0 || number > numsLen {
			nums[position] = 0
			position++
		} else {
			//1.本身呆在自己的位置的，无需任何操作
			if number == position+1 && number == nums[number-1] {
				position++
				//2.不在自己位置上的，但是自己位置上已经有正确数字的，自己本身置为0，指针下移
			} else if number != position+1 && number == nums[number-1] {
				nums[position] = 0
				position++
				//3.不在自己位置上，并且自己位置上没有正确数字，两个交换
			} else {
				nums[position] = nums[number-1]
				nums[number-1] = number
			}
		}
	}
	for i := 0; i < numsLen; i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return numsLen + 1
}

func testProblem41() {
	input1 := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(input1))
	input2 := []int{1}
	fmt.Println(firstMissingPositive(input2))
}
