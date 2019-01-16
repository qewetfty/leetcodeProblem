package array

import "fmt"

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	}
	i := 0
	j := numsLen - 1
	mid := 0
	left := -1
	right := -1
	// 先找target在最左侧的index
	for i <= j {
		mid = (i + j) / 2
		searchNum := nums[mid]
		if target == searchNum {
			if mid == 0 {
				left = 0
				break
			} else {
				if nums[mid-1] != target {
					left = mid
					break
				} else {
					j = mid - 1
				}
			}
		} else if searchNum < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	// 再找target最右侧的index
	i = 0
	j = numsLen - 1
	for i <= j {
		mid = (i + j) / 2
		searchNum := nums[mid]
		if target == searchNum {
			if mid == numsLen-1 {
				right = numsLen - 1
				break
			} else {
				if nums[mid+1] != target {
					right = mid
					break
				} else {
					i = mid + 1
				}
			}
		} else if searchNum < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return []int{left, right}
}

func testProblem34() {
	input1 := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(input1, 8))
	fmt.Println(searchRange(input1, 6))
	input2 := []int{1}
	fmt.Println(searchRange(input2, 1))
	input3 := []int{2, 2}
	fmt.Println(searchRange(input3, 3))
}
