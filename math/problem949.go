package math

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an array of 4 digits, return the largest 24 hour time that can be made.
// The smallest 24 hour time is 00:00, and the largest is 23:59.  Starting from 00:00, a time is larger if more time
// has elapsed since midnight.
// Return the answer as a string of length 5.  If no valid time can be made, return an empty string.
//	Example 1:
//		Input: [1,2,3,4]
//		Output: "23:41"
//	Example 2:
//		Input: [5,5,5,5]
//		Output: ""
//	Note:
//		A.length == 4
//		0 <= A[i] <= 9

func largestTimeFromDigits(A []int) string {
	ans := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if k == i || k == j {
					continue
				}
				hour, minute := 10*A[i]+A[j], 10*A[k]+A[6-i-j-k]
				if hour < 24 && minute < 60 {
					ans = utils.Max(ans, hour*60+minute)
				}
			}
		}
	}
	if ans < 0 {
		return ""
	}
	res := fmt.Sprintf("%02d:%02d", ans/60, ans%60)
	return res
}

func testProblem949() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	fmt.Println(largestTimeFromDigits([]int{5, 5, 5, 5}))
}
