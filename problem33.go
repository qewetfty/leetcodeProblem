package main

import (
	"fmt"
	"github.com/leetcodeProblem/main/array"
)

//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//You are given a target value to search. If found in the array return its index, otherwise return -1.
//You may assume no duplicate exists in the array.
//Your algorithm's runtime complexity must be in the order of O(log n).

func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	i := 0
	j := numsLen - 1
	first := nums[0]
	mid := 0
	for i <= j {
		mid = (i + j) / 2
		curNum := nums[mid]
		if target == first {
			return 0
		} else if target > first {
			if curNum == target {
				return mid
			}
			if curNum < first {
				j = mid - 1
			} else {
				if target < curNum {
					j = mid - 1
				} else {
					i = mid + 1
				}
			}
		} else {
			if curNum == target {
				return mid
			}
			if curNum >= first {
				i = mid + 1
			} else {
				if target > curNum {
					i = mid + 1
				} else {
					j = mid - 1
				}
			}
		}
	}
	return -1
}

func main() {
	input1 := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(input1, 0))
	fmt.Println(search(input1, 3))
	input2 := []int{3, 1}
	fmt.Println(search(input2, 1))
	input3 := []int{1, 3}
	fmt.Println(search(input3, 3))
	fmt.Println(array.SearchInsert(input1, 1))
}
