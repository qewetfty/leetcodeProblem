package array

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// Given an integer array nums, return the maximum difference between two
// successive elements in its sorted form. If the array contains less than two
// elements, return 0.
// You must write an algorithm that runs in linear time and uses linear extra space.
//	Example 1:
//		Input: nums = [3,6,9,1]
//		Output: 3
//		Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) has the maximum difference 3.
//	Example 2:
//		Input: nums = [10]
//		Output: 0
//		Explanation: The array contains less than 2 elements, therefore return 0.
//	Constraints:
//		1 <= nums.length <= 104
//		0 <= nums[i] <= 109

// 桶排序方法
// 将数据均分为l-1个桶，则桶之间的数据相邻差值为min-max。相邻数字之间的最大差值不会小于(max-min)/(l-1)
// 差值最大的数据一定是在相邻的桶
func maximumGap(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	maxVal, minVal := findMaxAndMinNum(nums)
	d := utils.Max(1, (maxVal-minVal)/(l-1))
	bucketNum := (maxVal-minVal)/d + 1
	buckets := make([]pair, bucketNum)
	for i := 0; i < bucketNum; i++ {
		buckets[i] = pair{
			max: -1,
			min: -1,
		}
	}
	for _, num := range nums {
		bucket := (num - minVal) / d
		if buckets[bucket].min == -1 {
			buckets[bucket].min, buckets[bucket].max = num, num
		} else {
			buckets[bucket].min, buckets[bucket].max = utils.Min(buckets[bucket].min, num), utils.Max(buckets[bucket].max, num)
		}
	}
	result := 0
	prev := -1
	for i, bucket := range buckets {
		if bucket.min == -1 {
			continue
		}
		if prev != -1 {
			result = utils.Max(result, bucket.min-buckets[prev].max)
		}
		prev = i
	}
	return result
}

func findMaxAndMinNum(nums []int) (int, int) {
	max, min := nums[0], nums[0]
	for _, num := range nums {
		max = utils.Max(max, num)
		min = utils.Min(min, num)
	}
	return max, min
}

type pair struct {
	max int
	min int
}

// 直接排序处理
func maximumGap2(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	sort.Ints(nums)
	result := nums[1] - nums[0]
	for i := 2; i < l; i++ {
		result = utils.Max(result, nums[i]-nums[i-1])
	}
	return result
}

func testProblem164() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
}
