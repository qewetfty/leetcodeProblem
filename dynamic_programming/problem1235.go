package dynamic_programming

import (
	"github.com/leetcodeProblem/utils"
	"sort"
)

// We have n jobs, where every job is scheduled to be done from startTime[i] to
// endTime[i], obtaining a profit of profit[i].
// You're given the startTime, endTime and profit arrays, return the maximum
// profit you can take such that there are no two jobs in the subset with
// overlapping time range.
// If you choose a job that ends at time X you will be able to start another job that starts at time X.
//	Example 1:
//		Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
//		Output: 120
//		Explanation: The subset chosen is the first and fourth job.
//					 Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.
//	Example 2:
//		Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
//		Output: 150
//		Explanation: The subset chosen is the first, fourth and fifth job.
//					 Profit obtained 150 = 20 + 70 + 60.
//	Example 3:
//		Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
//		Output: 6
//	Constraints:
//		1 <= startTime.length == endTime.length == profit.length <= 5 * 104
//		1 <= startTime[i] < endTime[i] <= 109
//		1 <= profit[i] <= 104

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	l := len(startTime)
	jobList := make([][]int, l)
	for i := 0; i < l; i++ {
		jobList[i] = make([]int, 3)
		jobList[i][0], jobList[i][1], jobList[i][2] = startTime[i], endTime[i], profit[i]
	}
	sort.Slice(jobList, func(i, j int) bool {
		return jobList[i][1] < jobList[j][1]
	})
	dp := make([]int, l)
	dp[0] = jobList[0][2]
	for i := 1; i < l; i++ {
		dp[i] = jobList[i][2]
		left, right := 0, i-1
		for left < right {
			mid := (left + right + 1) >> 1
			if jobList[mid][1] <= jobList[i][0] {
				left = mid
			} else {
				right = mid - 1
			}
		}
		if jobList[left][1] <= jobList[i][0] {
			dp[i] += dp[left]
		}
		dp[i] = utils.Max(dp[i-1], dp[i])
	}
	return dp[l-1]
}
