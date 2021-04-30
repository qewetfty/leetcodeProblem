package hashTable

import "fmt"

// Given three integers x, y, and bound, return a list of all the powerful
// integers that have a value less than or equal to bound.
// An integer is powerful if it can be represented as xi + yj for some integers i >= 0 and j >= 0.
// You may return the answer in any order. In your answer, each value should occur at most once.
//	Example 1:
//		Input: x = 2, y = 3, bound = 10
//		Output: [2,3,4,5,7,9,10]
//		Explanation:
//			2 = 20 + 30
//			3 = 21 + 30
//			4 = 20 + 31
//			5 = 21 + 31
//			7 = 22 + 31
//			9 = 23 + 30
//			10 = 20 + 32
//	Example 2:
//		Input: x = 3, y = 5, bound = 15
//		Output: [2,4,6,8,10,14]
//	Constraints:
//		1 <= x, y <= 100
//		0 <= bound <= 106

func powerfulIntegers(x int, y int, bound int) []int {
	xPowerList, yPowerList := make([]int, 0), make([]int, 0)
	if x == 1 {
		xPowerList = append(xPowerList, 1)
	} else {
		start := 1
		for start <= bound {
			xPowerList = append(xPowerList, start)
			start *= x
		}
	}
	if y == 1 {
		yPowerList = append(yPowerList, 1)
	} else {
		start := 1
		for start <= bound {
			yPowerList = append(yPowerList, start)
			start *= y
		}
	}

	resultMap := make(map[int]bool)
	for _, xPower := range xPowerList {
		for _, yPower := range yPowerList {
			answer := xPower + yPower
			if answer <= bound {
				resultMap[answer] = true
			}
		}
	}
	result := make([]int, 0)
	for key := range resultMap {
		result = append(result, key)
	}
	return result
}

func testProblem970() {
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(3, 5, 15))
	fmt.Println(powerfulIntegers(1, 4, 7))
	fmt.Println(powerfulIntegers(1, 1, 1000000))
}
