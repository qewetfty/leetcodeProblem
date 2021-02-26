package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// You are given two integer arrays nums and multipliers of size n and m respectively, where n >= m. The arrays are 1-indexed.
// You begin with a score of 0. You want to perform exactly m operations. On the ith operation (1-indexed), you will:
//	Choose one integer x from either the start or the end of the array nums.
//	Add multipliers[i] * x to your score.
//	Remove x from the array nums.
// Return the maximum score after performing m operations.
//	Example 1:
//		Input: nums = [1,2,3], multipliers = [3,2,1]
//		Output: 14
//		Explanation: An optimal solution is as follows:
//			- Choose from the end, [1,2,3], adding 3 * 3 = 9 to the score.
//			- Choose from the end, [1,2], adding 2 * 2 = 4 to the score.
//			- Choose from the end, [1], adding 1 * 1 = 1 to the score.
//			The total score is 9 + 4 + 1 = 14.
//	Example 2:
//		Input: nums = [-5,-3,-3,-2,7,1], multipliers = [-10,-5,3,4,6]
//		Output: 102
//		Explanation: An optimal solution is as follows:
//			- Choose from the start, [-5,-3,-3,-2,7,1], adding -5 * -10 = 50 to the score.
//			- Choose from the start, [-3,-3,-2,7,1], adding -3 * -5 = 15 to the score.
//			- Choose from the start, [-3,-2,7,1], adding -3 * 3 = -9 to the score.
//			- Choose from the end, [-2,7,1], adding 1 * 4 = 4 to the score.
//			- Choose from the end, [-2,7], adding 7 * 6 = 42 to the score.
//			The total score is 50 + 15 - 9 + 4 + 42 = 102.
//	Constraints:
//		n == nums.length
//		m == multipliers.length
//		1 <= m <= 103
//		m <= n <= 105
//		-1000 <= nums[i], multipliers[i] <= 1000

func maximumScore(nums []int, multipliers []int) int {
	m, n := len(nums), len(multipliers)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + nums[i-1]*multipliers[i-1]
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + nums[m-j]*multipliers[j-1]
	}
	for i := 1; i <= n; i++ {
		for j := 1; i+j <= n; j++ {
			dp[i][j] = utils.Max(dp[i-1][j]+nums[i-1]*multipliers[i+j-1], dp[i][j-1]+nums[m-j]*multipliers[i+j-1])
		}
	}
	result := math.MinInt32
	for i := 0; i <= n; i++ {
		result = utils.Max(result, dp[i][n-i])
	}
	return result
}

func testProblem1770() {
	fmt.Println(maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}))
	fmt.Println(maximumScore([]int{1, 2, 3}, []int{3, 2, 1}))
}
