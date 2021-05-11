package sliding_window

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// There are several cards arranged in a row, and each card has an associated
// number of points The points are given in the integer array cardPoints.
// In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
// Your score is the sum of the points of the cards you have taken.
// Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
//	Example 1:
//		Input: cardPoints = [1,2,3,4,5,6,1], k = 3
//		Output: 12
//		Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
//	Example 2:
//		Input: cardPoints = [2,2,2], k = 2
//		Output: 4
//		Explanation: Regardless of which two cards you take, your score will always be 4.
//	Example 3:
//		Input: cardPoints = [9,7,7,9,7,7,9], k = 7
//		Output: 55
//		Explanation: You have to take all the cards. Your score is the sum of points of all cards.
//	Example 4:
//		Input: cardPoints = [1,1000,1], k = 1
//		Output: 1
//		Explanation: You cannot take the card in the middle. Your best score is 1.
//	Example 5:
//		Input: cardPoints = [1,79,80,1,1,1,200,1], k = 3
//		Output: 202
//	Constraints:
//		1 <= cardPoints.length <= 10^5
//		1 <= cardPoints[i] <= 10^4
//		1 <= k <= cardPoints.length

func maxScore(cardPoints []int, k int) int {
	l := len(cardPoints)
	remain := l - k
	sum, windowSum := 0, 0
	for i, point := range cardPoints {
		if i < remain {
			windowSum += point
		}
		sum += point
	}
	result := sum - windowSum
	for i := remain; i < l; i++ {
		windowSum = windowSum - cardPoints[i-remain] + cardPoints[i]
		result = utils.Max(result, sum-windowSum)
	}
	return result
}

func testProblem1423() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
	fmt.Println(maxScore([]int{1, 1000, 1}, 1))
	fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
}
