package treeNode

import "github.com/leetcodeProblem/data"

// You are given the root of a binary search tree (BST), where exactly two nodes
// of the tree were swapped by mistake. Recover the tree without changing its
// structure.
// Follow up: A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?
//	Example 1:
//		Input: root = [1,3,null,null,2]
//		Output: [3,1,null,null,2]
//		Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.
//	Example 2:
//		Input: root = [3,1,4,null,null,2]
//		Output: [2,1,4,null,null,3]
//		Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.
//	Constraints:
//		The number of nodes in the tree is in the range [2, 1000].
//		-231 <= Node.val <= 231 - 1

// 莫里斯遍历方法，将栈的储存方式修改为使用当前节点的前驱节点的右子树保存，省去栈的存储空间。时间复杂度O(2n)=O(n)，空间复杂度O(1)
func recoverTree(root *data.TreeNode) {
	var (
		x, y        *data.TreeNode // 记录需要交换位置的两个节点
		prev        *data.TreeNode // 记录当前遍历节点的前一个节点
		predecessor *data.TreeNode // 记录当前节点的前驱节点
	)
	for root != nil {
		// 如果左子树不为空的话，需要先找到前驱节点，然后把前驱节点的右子树指向root，并继续遍历节点的左子树
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			// predecessor指向root，然后继续遍历左子树
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else { // 说明是已经连接到了，说明这次我们需要比较root了，然后需要把链接断开了
				if prev != nil && prev.Val > root.Val {
					y = root
					if x == nil {
						x = prev
					}
				}
				prev = root
				// 恢复树
				predecessor.Right = nil
				root = root.Right
			}
			// 如果左节点是空的话，则需要比较root了，然后继续遍历root的右子树
		} else {
			if prev != nil && prev.Val > root.Val {
				y = root
				if x == nil {
					x = prev
				}
			}
			prev = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

// 隐式中序遍历方法，仅比较相邻的两个数即可，无需使用list保存处理。时间复杂度O(n)，空间复杂度O(H)，H为二叉树高度。
func recoverTreeWithoutList(root *data.TreeNode) {
	var (
		x, y *data.TreeNode // 记录需要交换位置的两个节点
		prev *data.TreeNode // 记录当前遍历的节点的前一个节点，用于比较值，判断是否是错误的节点
	)
	// 使用一个栈对遍历的顺序进行保存
	stack := make([]*data.TreeNode, 0)
	for len(stack) > 0 || root != nil {
		// 遍历左子树，并入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出最顶的节点，开始进行比较
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val > root.Val {
			y = root
			if x == nil {
				x = prev
			} else {
				break
			}
		}
		// 遍历弹出节点的右子树
		prev, root = root, root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

// 显式的中序遍历方法，找到两个需要交换的点，然后进行交换。时间复杂度O(n)，空间复杂度O(n)
func recoverTreeWithList(root *data.TreeNode) {
	// 1.将中序遍历结果保存至列表中
	inorderResult := make([]*data.TreeNode, 0)
	var inorder func(root *data.TreeNode)
	inorder = func(root *data.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		inorderResult = append(inorderResult, root)
		inorder(root.Right)
	}
	inorder(root)
	l := len(inorderResult)
	// 2.寻找列表序列中排序不正确的位置，可能是有一个左比右大的情况，这样则交换相邻的两个节点；可能是两个左比右大的情况，这样则交换两个节点。
	var findSwapRoot func() (*data.TreeNode, *data.TreeNode)
	findSwapRoot = func() (*data.TreeNode, *data.TreeNode) {
		var x, y *data.TreeNode
		for i := 0; i < l-1; i++ {
			if inorderResult[i].Val > inorderResult[i+1].Val {
				y = inorderResult[i+1]
				if x == nil {
					x = inorderResult[i]
				} else {
					break
				}
			}
		}
		return x, y
	}
	x, y := findSwapRoot()
	// 3.交换节点值
	x.Val, y.Val = y.Val, x.Val
}
