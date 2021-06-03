package array

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// Given a rectangular cake with height h and width w, and two arrays of integers
// horizontalCuts and verticalCuts where horizontalCuts[i] is the distance from
// the top of the rectangular cake to the ith horizontal cut and similarly,
// verticalCuts[j] is the distance from the left of the rectangular cake to the
// jth vertical cut.
// Return the maximum area of a piece of cake after you cut at each horizontal
// and vertical position provided in the arrays horizontalCuts and verticalCuts.
// Since the answer can be a huge number, return this modulo 10^9 + 7.
//	Example 1:
//		Input: h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
//		Output: 4
//		Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green piece of cake has the maximum area.
//	Example 2:
//		Input: h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
//		Output: 6
//		Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green and yellow pieces of cake have the maximum area.
//	Example 3:
//		Input: h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
//		Output: 9
//	Constraints:
//		2 <= h, w <= 10^9
//		1 <= horizontalCuts.length < min(h, 10^5)
//		1 <= verticalCuts.length < min(w, 10^5)
//		1 <= horizontalCuts[i] < h
//		1 <= verticalCuts[i] < w
//		It is guaranteed that all elements in horizontalCuts are distinct.
//		It is guaranteed that all elements in verticalCuts are distinct.

func maxArea1465(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	m, n := len(horizontalCuts), len(verticalCuts)
	maxH, maxV := utils.Max(horizontalCuts[0]-0, h-horizontalCuts[m-1]), utils.Max(verticalCuts[0]-0, w-verticalCuts[n-1])
	for i := 1; i < m; i++ {
		maxH = utils.Max(maxH, horizontalCuts[i]-horizontalCuts[i-1])
	}
	for i := 1; i < n; i++ {
		maxV = utils.Max(maxV, verticalCuts[i]-verticalCuts[i-1])
	}
	return (maxH * maxV) % 1000000007
}

func testProblem1465() {
	fmt.Println(maxArea1465(5, 4, []int{1, 2, 4}, []int{1, 3}))
	fmt.Println(maxArea1465(5, 4, []int{3, 1}, []int{1}))
	fmt.Println(maxArea1465(5, 4, []int{3}, []int{3}))
}
