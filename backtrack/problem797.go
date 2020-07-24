package backtrack

import "fmt"

// Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1, and return them in any order.
// The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.  graph[i] is a list of all nodes j
// for which the edge (i, j) exists.
//	Example:
//		Input: [[1,2], [3], [3], []]
//		Output: [[0,1,3],[0,2,3]]
//	Explanation: The graph looks like this:
//		0--->1
//		|    |
//		v    v
//		2--->3
//		There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
//	Note:
//		The number of nodes in the graph will be in the range [2, 15].
//		You can print different paths in any order, but you should keep the order of nodes inside one path.

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	curResult := make([]int, 0)
	curResult = append(curResult, 0)
	backtrack797(graph, 0, curResult, &res)
	return res
}

func backtrack797(graph [][]int, curPos int, curResult []int, res *[][]int) {
	if curPos == len(graph)-1 {
		newResult := make([]int, 0)
		newResult = append(newResult, curResult...)
		*res = append(*res, newResult)
		return
	}
	for _, node := range graph[curPos] {
		curResult = append(curResult, node)
		backtrack797(graph, node, curResult, res)
		curResult = curResult[:len(curResult)-1]
	}
}

func testProblem797() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
