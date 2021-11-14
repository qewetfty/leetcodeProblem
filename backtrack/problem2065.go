package backtrack

import "leetcodeProblem/utils"

// There is an undirected graph with n nodes numbered from 0 to n - 1
// (inclusive). You are given a 0-indexed integer array values where values[i] is
// the value of the ith node. You are also given a 0-indexed 2D integer array
// edges, where each edges[j] = [uj, vj, timej] indicates that there is an
// undirected edge between the nodes uj and vj, and it takes timej seconds to
// travel between the two nodes. Finally, you are given an integer maxTime.
// A valid path in the graph is any path that starts at node 0, ends at node 0,
// and takes at most maxTime seconds to complete. You may visit the same node
// multiple times. The quality of a valid path is the sum of the values of the
// unique nodes visited in the path (each node's value is added at most once to
// the sum).
// Return the maximum quality of a valid path.
//	Note: There are at most four edges connected to each node.
//	Example 1:
//		Input: values = [0,32,10,43], edges = [[0,1,10],[1,2,15],[0,3,10]], maxTime = 49
//		Output: 75
//		Explanation:
//			One possible path is 0 -> 1 -> 0 -> 3 -> 0. The total time taken is 10 + 10 + 10 + 10 = 40 <= 49.
//			The nodes visited are 0, 1, and 3, giving a maximal path quality of 0 + 32 + 43 = 75.
//	Example 2:
//		Input: values = [5,10,15,20], edges = [[0,1,10],[1,2,10],[0,3,10]], maxTime = 30
//		Output: 25
//		Explanation:
//			One possible path is 0 -> 3 -> 0. The total time taken is 10 + 10 = 20 <= 30.
//			The nodes visited are 0 and 3, giving a maximal path quality of 5 + 20 = 25.
//	Example 3:
//		Input: values = [1,2,3,4], edges = [[0,1,10],[1,2,11],[2,3,12],[1,3,13]], maxTime = 50
//		Output: 7
//		Explanation:
//			One possible path is 0 -> 1 -> 3 -> 1 -> 0. The total time taken is 10 + 13 + 13 + 10 = 46 <= 50.
//			The nodes visited are 0, 1, and 3, giving a maximal path quality of 1 + 2 + 4 = 7.
//	Example 4:
//		Input: values = [0,1,2], edges = [[1,2,10]], maxTime = 10
//		Output: 0
//		Explanation:
//			The only path is 0. The total time taken is 0.
//			The only node visited is 0, giving a maximal path quality of 0.
//	Constraints:
//		n == values.length
//		1 <= n <= 1000
//		0 <= values[i] <= 108
//		0 <= edges.length <= 2000
//		edges[j].length == 3
//		0 <= uj < vj <= n - 1
//		10 <= timej, maxTime <= 100
//		All the pairs [uj, vj] are unique.
//		There are at most four edges connected to each node.
//		The graph may not be connected.

var visitedMap map[int]int

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	nodeMap := make(map[int][][]int)
	for _, edge := range edges {
		node1, node2, time := edge[0], edge[1], edge[2]
		nodeMap[node1] = append(nodeMap[node1], []int{node2, time})
		nodeMap[node2] = append(nodeMap[node2], []int{node1, time})
	}
	if len(nodeMap[0]) == 0 {
		return values[0]
	}
	// 判断经过了几次节点的map，初始在0节点上，因此0节点值为1
	visitedMap = make(map[int]int)
	visitedMap[0] = 1
	maxValue, useTime := values[0], 0
	backtrack2065(0, values, nodeMap, useTime, maxTime, &maxValue)
	return maxValue
}

func backtrack2065(node int, values []int, nodeMap map[int][][]int, useTime, maxTime int, maxValue *int) {
	if useTime > maxTime {
		return
	}
	// 遍历回了0节点，直接计算最后的结果
	if node == 0 {
		curValue := 0
		for curNode, visitTime := range visitedMap {
			if visitTime > 0 {
				curValue = curValue + values[curNode]
			}
		}
		*maxValue = utils.Max(*maxValue, curValue)
	}
	// 遍历节点连接的其他节点
	for _, next := range nodeMap[node] {
		nextNode, time := next[0], next[1]
		useTime, visitedMap[nextNode] = useTime+time, visitedMap[nextNode]+1
		backtrack2065(nextNode, values, nodeMap, useTime, maxTime, maxValue)
		useTime, visitedMap[nextNode] = useTime-time, visitedMap[nextNode]-1
	}
}
