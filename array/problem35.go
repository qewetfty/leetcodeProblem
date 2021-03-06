package array

//Given a sorted array and a target value, return the index
// if the target is found. If not, return the index where it
// would be if it were inserted in order.
//
//You may assume no duplicates in the array.

func SearchInsert(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	i := 0
	j := numsLen
	mid := 0
	for i <= j {
		mid = (i + j) / 2
		midNumber := nums[mid]
		if midNumber == target {
			return mid
		}
		if midNumber < target {
			if mid == numsLen-1 {
				return numsLen
			}
			if target < nums[mid+1] {
				return mid + 1
			}
			i = mid + 1
		} else {
			if mid == 0 {
				return 0
			}
			if nums[mid-1] < target {
				return mid
			}
			j = mid - 1
		}
	}
	return mid
}
