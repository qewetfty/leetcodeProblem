package dfs_bfs

import (
	"fmt"
	"math"
)

// A tree is an undirected graph in which any two vertices are connected by
// exactly one path. In other words, any connected graph without simple cycles is
// a tree.
//
// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges
// where edges[i] = [ai, bi] indicates that there is an undirected edge between
// the two nodes ai and bi in the tree, you can choose any node of the tree as
// the root. When you select a node x as the root, the result tree has height h.
// Among all possible rooted trees, those with minimum height (i.e. min(h)) are
// called minimum height trees (MHTs).
//
// Return a list of all MHTs' root labels. You can return the answer in any order.
// The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.
//	Example 1:
//		Input: n = 4, edges = [[1,0],[1,2],[1,3]]
//		Output: [1]
//		Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.
//	Example 2:
//		Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
//		Output: [3,4]
//	Example 3:
//		Input: n = 1, edges = []
//		Output: [0]
//	Example 4:
//		Input: n = 2, edges = [[0,1]]
//		Output: [0,1]
//	Constraints:
//		1 <= n <= 2 * 104
//		edges.length == n - 1
//		0 <= ai, bi < n
//		ai != bi
//		All the pairs (ai, bi) are distinct.
//		The given input is guaranteed to be a tree and there will be no repeated edges.

// 拓扑排序，时间复杂度O(n)
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}
	// 构建每个节点所连接的节点集合
	neighbors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]bool, 0)
	}
	for _, edge := range edges {
		neighbors[edge[0]][edge[1]], neighbors[edge[1]][edge[0]] = true, true
	}
	leaves := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(neighbors[i]) == 1 {
			leaves = append(leaves, i)
		}
	}
	remaining := n
	for remaining > 2 {
		remaining -= len(leaves)
		newLeaves := make([]int, 0)
		for _, leaf := range leaves {
			for neighbor := range neighbors[leaf] {
				delete(neighbors[neighbor], leaf)
				if len(neighbors[neighbor]) == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}
		leaves = newLeaves
	}
	return leaves
}

// bfs遍历每一个节点，时间复杂度O(n^2)，超时
func findMinHeightTrees2(n int, edges [][]int) []int {
	l := len(edges)
	if n == 1 || l == 0 {
		return []int{0}
	}
	// 构建每个节点所连接的节点集合
	nodeList := make([][]int, n)
	for i := 0; i < n; i++ {
		nodeList[i] = make([]int, 0)
	}
	for _, edge := range edges {
		nodeList[edge[0]], nodeList[edge[1]] = append(nodeList[edge[0]], edge[1]), append(nodeList[edge[1]], edge[0])
	}
	// bfs遍历每个结果
	mht, res := math.MaxInt64, make([]int, 0)
	for i := 0; i < n; i++ {
		queue, visited, height := make([]int, 0), make(map[int]bool), 0
		queue, visited[i] = append(queue, i), true
		for len(queue) > 0 {
			curLen := len(queue)
			height++
			for j := 0; j < curLen; j++ {
				curNode := queue[j]
				for _, node := range nodeList[curNode] {
					if !visited[node] {
						queue, visited[node] = append(queue, node), true
					}
				}
			}
			queue = queue[curLen:]
		}
		if height < mht {
			mht, res = height, make([]int, 0)
			res = append(res, i)
		} else if height == mht {
			res = append(res, i)
		}
	}
	return res
}

func testProblem310() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	fmt.Println(findMinHeightTrees(1, [][]int{}))
	fmt.Println(findMinHeightTrees(2, [][]int{{0, 1}}))
}
