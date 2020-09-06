package array

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Two images A and B are given, represented as binary, square matrices of the same size.  (A binary matrix has only 0s and 1s as values.)
// We translate one image however we choose (sliding it left, right, up, or down any number of units), and place it on
// top of the other image.  After, the overlap of this translation is the number of positions that have a 1 in both images.
// (Note also that a translation does not include any kind of rotation.)
// What is the largest possible overlap?
//	Example 1:
//		Input: A = [[1,1,0],
//		            [0,1,0],
//		            [0,1,0]]
//		       B = [[0,0,0],
//		            [0,1,1],
//		            [0,0,1]]
//		Output: 3
//		Explanation: We slide A to right by 1 unit and down by 1 unit.
//	Notes:
//		1 <= A.length = A[0].length = B.length = B[0].length <= 30
//		0 <= A[i][j], B[i][j] <= 1

func largestOverlap(A [][]int, B [][]int) int {
	N := len(A)
	delta := make([][]int, 2*N+1)
	for i := 0; i < 2*N+1; i++ {
		delta[i] = make([]int, 2*N+1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if A[i][j] == 1 {
				for k := 0; k < N; k++ {
					for l := 0; l < N; l++ {
						if B[k][l] == 1 {
							delta[i-k+N][j-l+N]++
						}
					}
				}
			}
		}
	}
	res := 0
	for _, nums := range delta {
		for _, num := range nums {
			res = utils.Max(res, num)
		}
	}
	return res
}

func testProblem835() {
	fmt.Println(largestOverlap([][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}))
}
