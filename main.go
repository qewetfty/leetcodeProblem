package main

import "fmt"

func findPairs2(nums []int, k int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	res := 0
	if k == 0 {
		for _, v := range numMap {
			if v > 1 {
				res++
			}
		}
	} else {
		for num := range numMap {
			if _, ok := numMap[num+k]; ok {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findPairs2([]int{1, 3, 1, 5, 4}, 0))
	fmt.Println(findPairs2([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs2([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs2([]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3))
	fmt.Println(findPairs2([]int{-1, -2, -3}, 1))
}
