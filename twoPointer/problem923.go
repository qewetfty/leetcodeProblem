package twoPointer

import (
	"fmt"
	"sort"
)

// Given an integer array arr, and an integer target, return the number of tuples
// i, j, k such that i < j < k and arr[i] + arr[j] + arr[k] == target.
// As the answer can be very large, return it modulo 109 + 7.
//	Example 1:
//		Input: arr = [1,1,2,2,3,3,4,4,5,5], target = 8
//		Output: 20
//		Explanation:
//			Enumerating by the values (arr[i], arr[j], arr[k]):
//			(1, 2, 5) occurs 8 times;
//			(1, 3, 4) occurs 8 times;
//			(2, 2, 4) occurs 2 times;
//			(2, 3, 3) occurs 2 times.
//	Example 2:
//		Input: arr = [1,1,2,2,2,2], target = 5
//		Output: 12
//		Explanation:
//			arr[i] = 1, arr[j] = arr[k] = 2 occurs 12 times:
//			We choose one 1 from [1,1] in 2 ways,
//			and two 2s from [2,2,2,2] in 6 ways.
//	Constraints:
//		3 <= arr.length <= 3000
//		0 <= arr[i] <= 100
//		0 <= target <= 300

var decimal = 1000000007

func threeSumMulti(arr []int, target int) int {
	numberMap, numberArr := make(map[int]int), make([]int, 0)
	for _, i := range arr {
		if _, ok := numberMap[i]; !ok {
			numberArr = append(numberArr, i)
		}
		numberMap[i]++
	}
	sort.Ints(numberArr)
	result := 0
	for i, num := range numberArr {
		twoSum := target - num
		if twoSum < num {
			break
		}
		lo, hi := i, len(numberArr)-1
		for lo <= hi {
			curTwoSum := numberArr[lo] + numberArr[hi]
			if curTwoSum == twoSum {
				if lo == i {
					// 三个数相等，计算组合数
					if lo == hi {
						if numberMap[numberArr[lo]] < 3 {
							lo++
							continue
						}
						count := numberMap[numberArr[lo]] * (numberMap[numberArr[lo]] - 1) * (numberMap[numberArr[lo]] - 2) / 6
						result = (result + count) % decimal
						// 前两个数相等
					} else {
						if numberMap[numberArr[lo]] < 2 {
							lo++
							continue
						}
						// 组合数计算
						count := numberMap[numberArr[lo]] * (numberMap[numberArr[lo]] - 1) / 2 * numberMap[numberArr[hi]]
						result = (result + count) % decimal
					}
				} else {
					// 后两个数相等
					if lo == hi {
						if numberMap[numberArr[lo]] < 2 {
							lo++
							continue
						}
						count := numberMap[numberArr[lo]] * (numberMap[numberArr[lo]] - 1) / 2 * numberMap[num]
						result = (result + count) % decimal
					} else {
						count := numberMap[num] * numberMap[numberArr[lo]] * numberMap[numberArr[hi]]
						result = (result + count) % decimal
					}
				}
				lo++
				hi--
			} else if curTwoSum < twoSum {
				lo++
			} else {
				hi--
			}
		}
	}
	return result
}

func testProblem923() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8}, 12))
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 2, 2}, 5))
}
