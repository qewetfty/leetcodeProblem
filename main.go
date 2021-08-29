package main

import (
	"github.com/leetcodeProblem/utils"
	"sort"
)

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

func main() {

}
