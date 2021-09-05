package dfs_bfs

// There is an undirected connected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.
// You are given the integer n and the array edges where edges[i] = [ai, bi]
// indicates that there is an edge between nodes ai and bi in the tree.
// Return an array answer of length n where answer[i] is the sum of the distances
// between the ith node in the tree and all other nodes.
//	Example 1:
//		Input: n = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
//		Output: [8,12,6,10,10,10]
//		Explanation: The tree is shown above.
//					 We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
//					 equals 1 + 1 + 2 + 2 + 2 = 8.
//					 Hence, answer[0] = 8, and so on.
//	Example 2:
//		Input: n = 1, edges = []
//		Output: [0]
//	Example 3:
//		Input: n = 2, edges = [[1,0]]
//		Output: [1,1]
//	Constraints:
//		1 <= n <= 3 * 104
//		edges.length == n - 1
//		edges[i].length == 2
//		0 <= ai, bi < n
//		ai != bi
//		The given input represents a valid tree.

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	dp, sz := make([]int, n), make([]int, n)
	var dfs func(u, f int)
	dfs = func(u, f int) {
		sz[u] = 1
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			dfs(v, u)
			dp[u] += dp[v] + sz[v]
			sz[u] += sz[v]
		}
	}
	dfs(0, -1)

	res := make([]int, n)
	var dfs2 func(u, f int)
	dfs2 = func(u, f int) {
		res[u] = dp[u]
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			pu, pv := dp[u], dp[v]
			su, sv := sz[u], sz[v]

			dp[u] -= dp[v] + sz[v]
			sz[u] -= sz[v]
			dp[v] += dp[u] + sz[u]
			sz[v] += sz[u]

			dfs2(v, u)
			dp[u], dp[v] = pu, pv
			sz[u], sz[v] = su, sv
		}
	}
	dfs2(0, -1)
	return res
}
