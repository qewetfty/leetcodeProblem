package main

import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	numberMap := make(map[int]int)
	for _, num := range changed {
		numberMap[num]++
	}
	result := make([]int, 0)
	for _, num := range changed {
		if value := numberMap[num]; value == 0 {
			continue
		}
		numberMap[num]--
		doubleNum := num * 2
		if value := numberMap[doubleNum]; value == 0 {
			return []int{}
		}
		numberMap[doubleNum]--
		result = append(result, num)
	}
	return result
}

func main() {
	fmt.Println(findOriginalArray([]int{1, 3, 4, 2, 6, 8}))
}
