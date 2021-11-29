package binarySearch

//Given a sorted array of distinct integers and a target value, return the
//index if the target is found. If not, return the index where it would be if it were
//inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
// Input: nums = [1,3,5,6], target = 5
//Output: 2
// Example 2:
// Input: nums = [1,3,5,6], target = 2
//Output: 1
// Example 3:
// Input: nums = [1,3,5,6], target = 7
//Output: 4
// Example 4:
// Input: nums = [1,3,5,6], target = 0
//Output: 0
// Example 5:
// Input: nums = [1], target = 0
//Output: 0
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums contains distinct values sorted in ascending order.
// -10⁴ <= target <= 10⁴
//
// Related Topics Array Binary Search 👍 5492 👎 335

func searchInsert(nums []int, target int) int {
	l := len(nums)
	lo, hi := 0, l-1
	result := 0
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			result = mid
			break
		} else if nums[mid] > target {
			result = mid
			hi--
		} else {
			result = mid + 1
			lo++
		}
	}
	return result
}
