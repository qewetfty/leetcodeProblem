package binarySearch

//Koko loves to eat bananas. There are n piles of bananas, the iᵗʰ pile has
//piles[i] bananas. The guards have gone and will come back in h hours.
//
// Koko can decide her bananas-per-hour eating speed of k. Each hour, she
//chooses some pile of bananas and eats k bananas from that pile. If the pile has less
//than k bananas, she eats all of them instead and will not eat any more bananas
//during this hour.
//
// Koko likes to eat slowly but still wants to finish eating all the bananas
//before the guards return.
//
// Return the minimum integer k such that she can eat all the bananas within h
//hours.
//
//
// Example 1:
//
//
//Input: piles = [3,6,7,11], h = 8
//Output: 4
//
//
// Example 2:
//
//
//Input: piles = [30,11,23,4,20], h = 5
//Output: 30
//
//
// Example 3:
//
//
//Input: piles = [30,11,23,4,20], h = 6
//Output: 23
//
//
//
// Constraints:
//
//
// 1 <= piles.length <= 10⁴
// piles.length <= h <= 10⁹
// 1 <= piles[i] <= 10⁹
//
// Related Topics Array Binary Search 👍 2541 👎 135

//leetcode submit region begin(Prohibit modification and deletion)
func minEatingSpeed(piles []int, h int) int {
	min, max := 1, -1
	for _, pile := range piles {
		max = Max(pile, max)
	}
	result := max
	for min <= max {
		mid := (min + max) >> 1
		needHour := 0
		for _, pile := range piles {
			needHour = needHour + (pile-1)/mid + 1
		}
		if needHour > h {
			min = mid + 1
		} else {
			result = mid
			max = mid - 1
		}
	}
	return result
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
