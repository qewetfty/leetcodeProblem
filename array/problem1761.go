package array

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// You are given an undirected graph. You are given an integer n which is the
// number of nodes in the graph and an array edges, where each edges[i] = [ui,
// vi] indicates that there is an undirected edge between ui and vi.
// A connected trio is a set of three nodes where there is an edge between every pair of them.
// The degree of a connected trio is the number of edges where one endpoint is in the trio, and the other is not.
// Return the minimum degree of a connected trio in the graph, or -1 if the graph has no connected trios.
//	Example 1:
//		Input: n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
//		Output: 3
//		Explanation: There is exactly one trio, which is [1,2,3]. The edges that form its degree are bolded in the figure above.
//	Example 2:
//		Input: n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
//		Output: 0
//		Explanation: There are exactly three trios:
//			1) [1,4,3] with degree 0.
//			2) [2,5,6] with degree 2.
//			3) [5,6,7] with degree 2.
//	Constraints:
//		2 <= n <= 400
//		edges[i].length == 2
//		1 <= edges.length <= n * (n-1) / 2
//		1 <= ui, vi <= n
//		ui != vi
//		There are no repeated edges.

func minTrioDegree(n int, edges [][]int) int {
	degreeList, connectMap := make([]int, n), make([][]bool, n)
	for i := 0; i < n; i++ {
		connectMap[i] = make([]bool, n)
	}
	for _, edge := range edges {
		connectMap[edge[0]-1][edge[1]-1] = true
		connectMap[edge[1]-1][edge[0]-1] = true
		degreeList[edge[0]-1]++
		degreeList[edge[1]-1]++
	}
	result := math.MaxInt32
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if connectMap[i][j] && connectMap[j][k] && connectMap[k][i] {
					result = utils.Min(result, degreeList[i]+degreeList[j]+degreeList[k]-6)
				}
			}
		}
	}
	if result == math.MaxInt32 {
		result = -1
	}
	return result
}

func testProblem1761() {
	fmt.Println(minTrioDegree(6, [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}))
	fmt.Println(minTrioDegree(7, [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}))
}
