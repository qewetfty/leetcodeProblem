package hashTable

import (
	"fmt"
	"sort"
)

// Given an array of citations (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.
// According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at least
// h citations each, and the other N − h papers have no more than h citations each."
//	Example:
//		Input: citations = [3,0,6,1,5]
//		Output: 3
//		Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had
//		             received 3, 0, 6, 1, 5 citations respectively.
//		             Since the researcher has 3 papers with at least 3 citations each and the remaining
//		             two with no more than 3 citations each, her h-index is 3.
//	Note: If there are several possible values for h, the maximum one is taken as the h-index.

func hIndex(citations []int) int {
	l := len(citations)
	if l == 0 {
		return 0
	}
	sort.Ints(citations)
	res := 0
	for i := 1; i <= l; i++ {
		if citations[l-i] >= i {
			res = i
		} else {
			break
		}
	}
	return res
}

func testProblem274() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(hIndex([]int{0, 0, 0, 0, 0, 0, 0}))
}
