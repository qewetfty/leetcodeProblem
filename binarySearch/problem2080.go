package binarySearch

// Design a data structure to find the frequency of a given value in a given subarray.
// The frequency of a value in a subarray is the number of occurrences of that value in the subarray.
// Implement the RangeFreqQuery class:
//	RangeFreqQuery(int[] arr) Constructs an instance of the class with the given 0-indexed integer array arr.
//	int query(int left, int right, int value) Returns the frequency of value in the subarray arr[left...right].
//	A subarray is a contiguous sequence of elements within an array.
//	arr[left...right] denotes the subarray that contains the elements of nums
//	between indices left and right (inclusive).
//	Example 1:
//		Input
//			["RangeFreqQuery", "query", "query"]
//			[[[12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56]], [1, 2, 4], [0, 11, 33]]
//		Output
//			[null, 1, 2]
//		Explanation
//			RangeFreqQuery rangeFreqQuery = new RangeFreqQuery([12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56]);
//			rangeFreqQuery.query(1, 2, 4); // return 1. The value 4 occurs 1 time in the subarray [33, 4]
//			rangeFreqQuery.query(0, 11, 33); // return 2. The value 33 occurs 2 times in the whole array.
//	Constraints:
//		1 <= arr.length <= 105
//		1 <= arr[i], value <= 104
//		0 <= left <= right < arr.length
//		At most 105 calls will be made to query

type RangeFreqQuery struct {
	// 存储各个number出现的索引，从小到大排列
	indexMap map[int][]int
}

func Constructor2080(arr []int) RangeFreqQuery {
	indexMap := make(map[int][]int)
	for i, num := range arr {
		indexMap[num] = append(indexMap[num], i)
	}
	return RangeFreqQuery{
		indexMap: indexMap,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	indexList := this.indexMap[value]
	if len(indexList) == 0 {
		return 0
	}
	// left比最大的index大，或者right比最小的index小，则没有
	if left > indexList[len(indexList)-1] || right < indexList[0] {
		return 0
	}
	leftIndex := findLeft(indexList, left)
	rightIndex := findRight(indexList, right)
	return rightIndex - leftIndex + 1
}

// 返回索引号
func findLeft(nums []int, val int) int {
	lo, hi := 0, len(nums)-1
	result := hi
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == val {
			result = mid
			break
		} else if nums[mid] < val {
			lo = mid + 1
		} else {
			result = mid
			hi = mid - 1
		}
	}
	return result
}

// 返回索引号
func findRight(nums []int, val int) int {
	lo, hi := 0, len(nums)-1
	result := lo
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == val {
			result = mid
			break
		} else if nums[mid] < val {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}
