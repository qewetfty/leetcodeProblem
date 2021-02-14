package dfs_bfs

import "fmt"

// Given an undirected graph, return true if and only if it is bipartite.
// Recall that a graph is bipartite if we can split its set of nodes into two
// independent subsets A and B, such that every edge in the graph has one node in
// A and another node in B.
// The graph is given in the following form: graph[i] is a list of indexes j for
// which the edge between nodes i and j exists. Each node is an integer between 0
// and graph.length - 1. There are no self edges or parallel edges: graph[i] does
// not contain i, and it doesn't contain any element twice.
//	Example 1:
//		Input: graph = [[1,3],[0,2],[1,3],[0,2]]
//		Output: true
//		Explanation: We can divide the vertices into two groups: {0, 2} and {1, 3}.
//	Example 2:
//		Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
//		Output: false
//		Explanation: We cannot find a way to divide the set of nodes into two independent subsets.
//	Constraints:
//		1 <= graph.length <= 100
//		0 <= graph[i].length < 100
//		0 <= graph[i][j] <= graph.length - 1
//		graph[i][j] != i
//		All the values of graph[i] are unique.
//		The graph is guaranteed to be undirected.

// 染色法，从0号节点开始，将0号节点置为红色，与之相连的置为绿色。
// -- 0 未染色   -- 1 红色    -- 2 绿色
// 如果遍历完了没有染色重合，返回true
// 如果染色重合了返回false

// bfs遍历
func isBipartite(graph [][]int) bool {
	l := len(graph)
	color := make([]int, l)
	for i := 0; i < l; i++ {
		if color[i] == 0 {
			color[i] = 1
			queue := make([]int, 0)
			queue = append(queue, i)
			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]
				neighbor, nodeColor := graph[node], color[node]
				for _, neighborNode := range neighbor {
					if color[neighborNode] == 0 {
						color[neighborNode] = nodeColor%2 + 1
						queue = append(queue, neighborNode)
					} else if color[neighborNode] == nodeColor {
						return false
					}
				}
			}
		}
	}
	return true
}

func testProblem785() {
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}))
}
