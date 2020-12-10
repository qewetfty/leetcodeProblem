package stack_queue

import "github.com/leetcodeProblem/data"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):
// BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The
// root of the BST is given as part of the constructor. The pointer should be
// initialized to a non-existent number smaller than any element in the BST.
// boolean hasNext() Returns true if there exists a number in the traversal to
// the right of the pointer, otherwise returns false. int next() Moves the
// pointer to the right, then returns the number at the pointer.
// Notice that by initializing the pointer to a non-existent smallest number, the first call to
// next() will return the smallest element in the BST.
// You may assume that next() calls will always be valid. That is, there will be
// at least a next number in the in-order traversal when next() is called.
//	Example 1:
//		Input
//			["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
//			[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
//		Output
//			[null, 3, 7, true, 9, true, 15, true, 20, false]
//		Explanation
//			BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
//			bSTIterator.next();    // return 3
//			bSTIterator.next();    // return 7
//			bSTIterator.hasNext(); // return True
//			bSTIterator.next();    // return 9
//			bSTIterator.hasNext(); // return True
//			bSTIterator.next();    // return 15
//			bSTIterator.hasNext(); // return True
//			bSTIterator.next();    // return 20
//			bSTIterator.hasNext(); // return False
//	Constraints:
//		The number of nodes in the tree is in the range [1, 105].
//		0 <= Node.val <= 106
//		At most 105 calls will be made to hasNext, and next.
//	Follow up:
//		Could you implement next() and hasNext() to run in average O(1) time and use O(h) memory, where h is the height of the tree?

type BSTIterator struct {
	stack []*data.TreeNode
}

func Constructor173(root *data.TreeNode) BSTIterator {
	stack := make([]*data.TreeNode, 0)
	iterator := BSTIterator{stack: stack}
	iterator.helper(root)
	return iterator
}

func (this *BSTIterator) helper(node *data.TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if node.Right != nil {
		this.helper(node.Right)
	}
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// 中序遍历列表做法
type BSTIterator1 struct {
	nodeList []*data.TreeNode
}

func Constructor1(root *data.TreeNode) BSTIterator1 {
	nodeList := make([]*data.TreeNode, 0)
	preorder(root, &nodeList)
	return BSTIterator1{nodeList: nodeList}
}

func preorder(node *data.TreeNode, nodeList *[]*data.TreeNode) {
	if node == nil {
		return
	}
	preorder(node.Left, nodeList)
	*nodeList = append(*nodeList, node)
	preorder(node.Right, nodeList)
}

func (this *BSTIterator1) Next1() int {
	res := this.nodeList[0]
	this.nodeList = this.nodeList[1:]
	return res.Val
}

func (this *BSTIterator1) HasNext1() bool {
	return len(this.nodeList) > 0
}

/**
 * Your BSTIterator1 object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
