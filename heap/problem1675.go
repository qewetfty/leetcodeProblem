package heap

import (
	"container/heap"
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// You are given an array nums of n positive integers.
// You can perform two types of operations on any element of the array any number of times:
//		If the element is even, divide it by 2.
//		For example, if the array is [1,2,3,4], then you can do this operation on the last element, and the array will be [1,2,3,2].
//		If the element is odd, multiply it by 2.
//		For example, if the array is [1,2,3,4], then you can do this operation on the first element, and the array will be [2,2,3,4].
//		The deviation of the array is the maximum difference between any two elements in the array.
// Return the minimum deviation the array can have after performing some number of operations.
//	Example 1:
//		Input: nums = [1,2,3,4]
//		Output: 1
//		Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.
//	Example 2:
//		Input: nums = [4,1,5,20,3]
//		Output: 3
//		Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.
//	Example 3:
//		Input: nums = [2,10,8]
//		Output: 3
//	Constraints:
//		n == nums.length
//		2 <= n <= 105
//		1 <= nums[i] <= 109

func minimumDeviation(nums []int) int {
	intHeap := new(bigIntHeap)
	heap.Init(intHeap)
	min := math.MaxInt32
	for _, num := range nums {
		if num&1 == 1 {
			num <<= 1
		}
		heap.Push(intHeap, num)
		min = utils.Min(min, num)
	}
	result := math.MaxInt32
	for true {
		curNum := heap.Pop(intHeap).(int)
		result = utils.Min(result, curNum-min)
		if curNum&1 == 1 {
			break
		}
		min = utils.Min(min, curNum>>1)
		heap.Push(intHeap, curNum>>1)
	}
	return result
}

type bigIntHeap []int

func (b bigIntHeap) Len() int {
	return len(b)
}

func (b bigIntHeap) Less(i, j int) bool {
	return b[j] < b[i]
}

func (b bigIntHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b *bigIntHeap) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *bigIntHeap) Pop() interface{} {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[0 : n-1]
	return x
}

func testProblem1675() {
	fmt.Println(minimumDeviation([]int{2, 10, 8}))
	fmt.Println(minimumDeviation([]int{1, 2, 3, 4}))
	fmt.Println(minimumDeviation([]int{4, 1, 5, 20, 3}))
}
