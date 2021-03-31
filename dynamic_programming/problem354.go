package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi]
// represents the width and the height of an envelope.
// One envelope can fit into another if and only if both the width and height of
// one envelope is greater than the width and height of the other envelope.
// Return the maximum number of envelopes can you Russian doll (i.e., put one inside the other).
//	Note: You cannot rotate an envelope.
//	Example 1:
//		Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
//		Output: 3
//		Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
//	Example 2:
//		Input: envelopes = [[1,1],[1,1],[1,1]]
//		Output: 1
//	Constraints:
//		1 <= envelopes.length <= 5000
//		envelopes[i].length == 2
//		1 <= wi, hi <= 104

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] < envelopes[j][1]
	})
	l := len(envelopes)
	dp := make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		// 以envelopes[i]做为最外层时的最大套娃数
		curEnvelop, curMaxCount := envelopes[i], 1
		for j := i - 1; j >= 0; j-- {
			curCount := 1
			if curEnvelop[0] > envelopes[j][0] && curEnvelop[1] > envelopes[j][1] {
				curCount = dp[j] + 1
			}
			curMaxCount = utils.Max(curMaxCount, curCount)
		}
		dp[i] = curMaxCount
	}
	result := 0
	for _, count := range dp {
		result = utils.Max(result, count)
	}
	return result
}

func testProblem354() {
	fmt.Println(maxEnvelopes([][]int{{46, 89}, {50, 53}, {52, 68}, {72, 45}, {77, 81}}))
	fmt.Println(maxEnvelopes([][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}))
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}))
}
