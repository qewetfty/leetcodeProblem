package greedy

import "fmt"

// You have a long flowerbed in which some of the plots are planted, and some are
// not. However, flowers cannot be planted in adjacent plots.
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty
// and 1 means not empty, and an integer n, return if n new flowers can be
// planted in the flowerbed without violating the no-adjacent-flowers rule.
//	Example 1:
//		Input: flowerbed = [1,0,0,0,1], n = 1
//		Output: true
//	Example 2:
//		Input: flowerbed = [1,0,0,0,1], n = 2
//		Output: false
//	Constraints:
//		1 <= flowerbed.length <= 2 * 104
//		flowerbed[i] is 0 or 1.
//		There are no two adjacent flowers in flowerbed.
//		0 <= n <= flowerbed.length

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	canPlaceNumber, lastPlacePos := 0, -2
	for i, place := range flowerbed {
		// 如果没放花，则继续遍历
		if place == 0 {
			continue
		}
		distance := i - lastPlacePos - 1
		lastPlacePos = i
		curPlaceNum := (distance - 1) / 2
		canPlaceNumber += curPlaceNum
	}
	distance := l + 1 - lastPlacePos - 1
	curPlaceNum := (distance - 1) / 2
	canPlaceNumber += curPlaceNum
	return canPlaceNumber >= n
}

func testProblem605() {
	fmt.Println(canPlaceFlowers([]int{0, 0}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2))
	fmt.Println(canPlaceFlowers([]int{0}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}
