package unionFind

import (
	"leetcodeProblem/data"
	"sort"
)

//You are given a string s, and an array of pairs of indices in the string
//pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.
//
// You can swap the characters at any pair of indices in the given pairs any
//number of times.
//
// Return the lexicographically smallest string that s can be changed to after
//using the swaps.
//
//
// Example 1:
//
//
//Input: s = "dcab", pairs = [[0,3],[1,2]]
//Output: "bacd"
//Explaination:
//Swap s[0] and s[3], s = "bcad"
//Swap s[1] and s[2], s = "bacd"
//
//
// Example 2:
//
//
//Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
//Output: "abcd"
//Explaination:
//Swap s[0] and s[3], s = "bcad"
//Swap s[0] and s[2], s = "acbd"
//Swap s[1] and s[2], s = "abcd"
//
// Example 3:
//
//
//Input: s = "cba", pairs = [[0,1],[1,2]]
//Output: "abc"
//Explaination:
//Swap s[0] and s[1], s = "bca"
//Swap s[1] and s[2], s = "bac"
//Swap s[0] and s[1], s = "abc"
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10^5
// 0 <= pairs.length <= 10^5
// 0 <= pairs[i][0], pairs[i][1] < s.length
// s only contains lower case English letters.
//
// Related Topics Hash Table String Depth-First Search Breadth-First Search
//Union Find 👍 1545 👎 52

//leetcode submit region begin(Prohibit modification and deletion)
func smallestStringWithSwaps(s string, pairs [][]int) string {
	l := len(s)
	u := data.NewUnionFind(l)
	for _, pair := range pairs {
		u.Union(pair[0], pair[1])
	}
	// 记录每个组都有哪些位置
	group := make(map[int][]int)
	for i := 0; i < l; i++ {
		group[u.Parent(i)] = append(group[u.Parent(i)], i)
	}
	// 创建结果的byte数组，按分组进行排序
	resultByte := make([]byte, l)
	for _, position := range group {
		posBytes := make([]byte, len(position))
		for i, pos := range position {
			posBytes[i] = s[pos]
		}
		sort.Slice(posBytes, func(i, j int) bool {
			return posBytes[i] < posBytes[j]
		})
		for i, pos := range position {
			resultByte[pos] = posBytes[i]
		}
	}
	return string(resultByte)
}
