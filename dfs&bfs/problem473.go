package dfs_bfs

import "fmt"

// You are given an integer array matchsticks where matchsticks[i] is the length
// of the ith matchstick. You want to use all the matchsticks to make one square.
// You should not break any stick, but you can link them up, and each matchstick
// must be used exactly one time.
// Return true if you can make this square and false otherwise.
//	Example 1:
//		Input: matchsticks = [1,1,2,2,2]
//		Output: true
//		Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.
//	Example 2:
//		Input: matchsticks = [3,3,3,3,4]
//		Output: false
//		Explanation: You cannot find a way to form a square with all the matchsticks.
//	Constraints:
//		1 <= matchsticks.length <= 15
//		0 <= matchsticks[i] <= 109

var matchsticksCombine []int
var l473 int

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, matchstick := range matchsticks {
		sum += matchstick
	}
	if sum%4 != 0 {
		return false
	}
	avg := sum / 4
	matchsticksCombine, l473 = make([]int, 4), len(matchsticks)
	return dfs473(matchsticks, 0, avg)
}

func dfs473(matchsticks []int, curr, avg int) bool {
	if curr == l473 {
		return checkCanBeSquare(avg)
	}
	for j := 0; j < 4; j++ {
		if matchsticksCombine[j]+matchsticks[curr] <= avg {
			matchsticksCombine[j] += matchsticks[curr]
			if dfs473(matchsticks, curr+1, avg) {
				return true
			}
			matchsticksCombine[j] -= matchsticks[curr]
		}
	}

	return false
}

func checkCanBeSquare(avg int) bool {
	for _, c := range matchsticksCombine {
		if c != avg {
			return false
		}
	}
	return true
}

func testProblem473() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
	fmt.Println(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}))
}
