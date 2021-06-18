package binaryIndexedTree

// Given an integer array nums, handle multiple queries of the following types:
//	Update the value of an element in nums.
//	Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
// Implement the NumArray class:
//	NumArray(int[] nums) Initializes the object with the integer array nums.
//	void update(int index, int val) Updates the value of nums[index] to be val.
//	int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
//	Example 1:
//		Input
//			["NumArray", "sumRange", "update", "sumRange"]
//			[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//		Output
//			[null, 9, null, 8]
//		Explanation
//			NumArray numArray = new NumArray([1, 3, 5]);
//			numArray.sumRange(0, 2); // return 1 + 3 + 5 = 9
//			numArray.update(1, 2);   // nums = [1, 2, 5]
//			numArray.sumRange(0, 2); // return 1 + 2 + 5 = 8
//	Constraints:
//		1 <= nums.length <= 3 * 104
//		-100 <= nums[i] <= 100
//		0 <= index < nums.length
//		-100 <= val <= 100
//		0 <= left <= right < nums.length
//		At most 3 * 104 calls will be made to update and sumRange.

type NumArray struct {
	nums []int
	bit  *bit307
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	bit := &bit307{
		nums:   make([]int, l+1),
		length: l + 1,
	}
	for i := 1; i <= l; i++ {
		bit.update(i, nums[i-1])
	}
	return NumArray{
		nums: nums,
		bit:  bit,
	}
}

func (this *NumArray) Update(index int, val int) {
	changeNum := val - this.nums[index]
	this.nums[index] = val
	this.bit.update(index+1, changeNum)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.bit.getSum(right+1) - this.bit.getSum(left+1)
}

type bit307 struct {
	nums   []int
	length int
}

func (b *bit307) lowBit(i int) int {
	return i & (-i)
}

func (b *bit307) update(index int, add int) {
	for i := index; i < b.length; i += b.lowBit(i) {
		b.nums[i] += add
	}
}

func (b *bit307) getSum(index int) int {
	result := 0
	for i := index; i > 0; i -= b.lowBit(i) {
		result += b.nums[i]
	}
	return result
}
