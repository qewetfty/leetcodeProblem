package heap

import (
	"container/heap"
	"leetcodeProblem/utils"
	"math"
)

//You are given an integer n denoting the number of nodes of a weighted
//directed graph. The nodes are numbered from 0 to n - 1.
//
// You are also given a 2D integer array edges where edges[i] = [fromi, toi,
//weighti] denotes that there exists a directed edge from fromi to toi with weight
//weighti.
//
// Lastly, you are given three distinct integers src1, src2, and dest denoting
//three distinct nodes of the graph.
//
// Return the minimum weight of a subgraph of the graph such that it is
//possible to reach dest from both src1 and src2 via a set of edges of this subgraph. In
//case such a subgraph does not exist, return -1.
//
// A subgraph is a graph whose vertices and edges are subsets of the original
//graph. The weight of a subgraph is the sum of weights of its constituent edges.
//
//
// Example 1:
//
//
//Input: n = 6, edges = [[0,2,2],[0,5,6],[1,0,3],[1,4,5],[2,1,1],[2,3,3],[2,3,4]
//,[3,4,2],[4,5,1]], src1 = 0, src2 = 1, dest = 5
//Output: 9
//Explanation:
//The above figure represents the input graph.
//The blue edges represent one of the subgraphs that yield the optimal answer.
//Note that the subgraph [[1,0,3],[0,5,6]] also yields the optimal answer. It
//is not possible to get a subgraph with less weight satisfying all the constraints.
//
//
//
// Example 2:
//
//
//Input: n = 3, edges = [[0,1,1],[2,1,1]], src1 = 0, src2 = 1, dest = 2
//Output: -1
//Explanation:
//The above figure represents the input graph.
//It can be seen that there does not exist any path from node 1 to node 2,
//hence there are no subgraphs satisfying all the constraints.
//
//
//
// Constraints:
//
//
// 3 <= n <= 10⁵
// 0 <= edges.length <= 10⁵
// edges[i].length == 3
// 0 <= fromi, toi, src1, src2, dest <= n - 1
// fromi != toi
// src1, src2, and dest are pairwise distinct.
// 1 <= weight[i] <= 10⁵
//
// Related Topics Graph Shortest Path 👍 372 👎 11

//leetcode submit region begin(Prohibit modification and deletion)
func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	g := make([][]edge, n)
	rg := make([][]edge, n)
	for _, e := range edges {
		from, to, weight := e[0], e[1], e[2]
		g[from] = append(g[from], edge{to, weight})
		rg[to] = append(rg[to], edge{from, weight})
	}
	d1 := dijkstra(g, src1)
	d2 := dijkstra(g, src2)
	d3 := dijkstra(rg, dest)

	res := math.MaxInt64 / 3
	for i := 0; i < n; i++ {
		res = utils.Min(res, d1[i]+d2[i]+d3[i])
	}
	if res < math.MaxInt64/3 {
		return int64(res)
	}
	return -1
}

func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = math.MaxInt64 / 3
	}
	// 初始化开始节点，其到自己的距离为0
	dis[start] = 0
	h := &hp{{start, 0}}
	for h.Len() > 0 {
		node := heap.Pop(h).(pair)
		v := node.v
		// 若不存在最小距离，则继续
		if node.dis > dis[v] {
			continue
		}
		// 遍历相邻节点
		for _, e := range g[v] {
			w := e.to
			if newDis := dis[v] + e.weight; newDis < dis[w] {
				dis[w] = newDis
				heap.Push(h, pair{w, newDis})
			}
		}
	}
	return dis
}

// 有向图的边
type edge struct {
	to     int //目标节点
	weight int //权重
}

// 有向图中，某个节点到其他节点的距离
type pair struct {
	v   int //有向图中另外一个节点
	dis int //距离
}

// 有向图中某个点到其他节点的距离数组，继承heap，可以快速获取最短距离
type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].dis < h[j].dis
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() interface{} {
	a := *h
	x := a[len(a)-1]
	*h = a[:len(a)-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
