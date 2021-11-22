package binarySearch

import (
	"github.com/leetcodeProblem/utils"
	"sort"
)

//You are given a 2D integer array items where items[i] = [pricei, beautyi]
//denotes the price and beauty of an item respectively.
//
// You are also given a 0-indexed integer array queries. For each queries[j],
//you want to determine the maximum beauty of an item whose price is less than or
//equal to queries[j]. If no such item exists, then the answer to this query is 0.
//
// Return an array answer of the same length as queries where answer[j] is the
//answer to the jᵗʰ query.
//
//
// Example 1:
//
//
//Input: items = [[1,2],[3,2],[2,4],[5,6],[3,5]], queries = [1,2,3,4,5,6]
//Output: [2,4,5,5,6,6]
//Explanation:
//- For queries[0]=1, [1,2] is the only item which has price <= 1. Hence, the
//answer for this query is 2.
//- For queries[1]=2, the items which can be considered are [1,2] and [2,4].
//  The maximum beauty among them is 4.
//- For queries[2]=3 and queries[3]=4, the items which can be considered are [1,
//2], [3,2], [2,4], and [3,5].
//  The maximum beauty among them is 5.
//- For queries[4]=5 and queries[5]=6, all items can be considered.
//  Hence, the answer for them is the maximum beauty of all items, i.e., 6.
//
//
// Example 2:
//
//
//Input: items = [[1,2],[1,2],[1,3],[1,4]], queries = [1]
//Output: [4]
//Explanation:
//The price of every item is equal to 1, so we choose the item with the maximum
//beauty 4.
//Note that multiple items can have the same price and/or beauty.
//
//
// Example 3:
//
//
//Input: items = [[10,1000]], queries = [5]
//Output: [0]
//Explanation:
//No item has a price less than or equal to 5, so no item can be chosen.
//Hence, the answer to the query is 0.
//
//
//
// Constraints:
//
//
// 1 <= items.length, queries.length <= 10⁵
// items[i].length == 2
// 1 <= pricei, beautyi, queries[j] <= 10⁹
//
// Related Topics Array Binary Search Sorting 👍 176 👎 3

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] != items[j][0] {
			return items[i][0] < items[j][0]
		}
		return items[i][1] > items[j][1]
	})
	// 构造当查询是某个值时，最大的结果
	l := len(items)
	maxList := make([]int, l)
	maxList[0] = items[0][1]
	for i := 1; i < l; i++ {
		maxList[i] = utils.Max(maxList[i-1], items[i][1])
	}
	// 二分法查找
	result := make([]int, 0)
	for _, query := range queries {
		index := binarySearch2070(items, query)
		if index == -1 {
			result = append(result, 0)
		} else {
			result = append(result, maxList[index])
		}
	}
	return result
}

func binarySearch2070(items [][]int, query int) int {
	lo, hi := 0, len(items)-1
	result := -1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if items[mid][0] == query {
			return mid
		} else if items[mid][0] < query {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}
