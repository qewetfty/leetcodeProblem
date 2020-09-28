package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	l, res, left, product := len(nums), 0, 0, 1
	for right := 0; right < l; right++ {
		product = product * nums[right]
		for product >= k {
			product = product / nums[left]
			left++
		}
		res = res + right - left + 1
	}
	return res
}

// time exceed method
func numSubarrayProductLessThanK2(nums []int, k int) int {
	l, res := len(nums), 0
	for i := 0; i < l; i++ {
		product := nums[i]
		if product >= k {
			continue
		} else {
			res++
		}
		for j := i + 1; j < l; j++ {
			product = product * nums[j]
			if product < k {
				res++
			} else {
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
