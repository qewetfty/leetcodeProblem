package dfs_bfs

import (
	"container/list"
	"github.com/leetcodeProblem/data"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// Given a binary tree, return the vertical order traversal of its nodes values.
// For each node at position (X, Y), its left and right children respectively will be at positions (X-1, Y-1) and (X+1, Y-1).
// Running a vertical line from X = -infinity to X = +infinity, whenever the vertical line touches some nodes,
// we report the values of the nodes in order from top to bottom (decreasing Y coordinates).
// If two nodes have the same position, then the value of the node that is reported first is the value that is smaller.
// Return an list of non-empty reports in order of X coordinate.  Every report will have a list of values of nodes.
//	Example 1:
//		Input: [3,9,20,null,null,15,7]
//		Output: [[9],[3,15],[20],[7]]
//		Explanation:
//			Without loss of generality, we can assume the root node is at position (0, 0):
//			Then, the node with value 9 occurs at position (-1, -1);
//			The nodes with values 3 and 15 occur at positions (0, 0) and (0, -2);
//			The node with value 20 occurs at position (1, -1);
//			The node with value 7 occurs at position (2, -2).
//	Example 2:
//		Input: [1,2,3,4,5,6,7]
//		Output: [[4],[2],[1,5,6],[3],[7]]
//		Explanation:
//			The node with value 5 and the node with value 6 have the same position according to the given scheme.
//			However, in the report "[1,5,6]", the node value of 5 comes first since 5 is smaller than 6.
//	Note:
//		The tree will have between 1 and 1000 nodes.
//		Each node's value will be between 0 and 1000.

// 相应的bfs遍历的数据结构，需要知道当前的x值和当前的节点
type Root struct {
	node *data.TreeNode
	x    int
}

func verticalTraversal(root *data.TreeNode) [][]int {
	res := make([][]int, 0)
	// 最终的每个x坐标对应的map
	resMap := make(map[int][]int)
	queue := new(list.List)
	// 记录x的最小、最大值，方便对结果进行遍历。
	lo, hi := 0, 0
	// 将开始的节点压入队列
	queue.PushFront(Root{
		node: root,
		x:    0,
	})
	for queue.Len() > 0 {
		curLen := queue.Len()
		// 当前层的x对应map，其中相同x的应按照大小升序排列，最后在进入最终x对应坐标的时候排序。
		curXMap := make(map[int][]int)
		for i := 0; i < curLen; i++ {
			curNode := queue.Remove(queue.Back()).(Root)
			node, x := curNode.node, curNode.x
			if _, ok := curXMap[x]; ok {
				curXMap[x] = append(curXMap[x], node.Val)
			} else {
				curXMap[x] = []int{node.Val}
			}
			if node.Left != nil {
				queue.PushFront(Root{
					node: node.Left,
					x:    x - 1,
				})
				lo = utils.Min(x-1, lo)
			}
			if node.Right != nil {
				queue.PushFront(Root{
					node: node.Right,
					x:    x + 1,
				})
				hi = utils.Max(x+1, hi)
			}
		}
		// 将当前层的结果加入到最终结果中
		for x, xnums := range curXMap {
			sort.Ints(xnums)
			if _, ok := resMap[x]; ok {
				resMap[x] = append(resMap[x], xnums...)
			} else {
				resMap[x] = xnums
			}
		}
	}
	// 写入结果
	for i := lo; i <= hi; i++ {
		res = append(res, resMap[i])
	}
	return res
}
