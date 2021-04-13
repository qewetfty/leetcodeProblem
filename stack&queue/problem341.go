package stack_queue

// You are given a nested list of integers nestedList. Each element is either an
// integer or a list whose elements may also be integers or other lists.
// Implement an iterator to flatten it.
//	Implement the NestedIterator class:
//		NestedIterator(List<NestedInteger> nestedList) Initializes the iterator with the nested list nestedList.
//		int next() Returns the next integer in the nested list.
//		boolean hasNext() Returns true if there are still some integers in the nested list and false otherwise.
//	Example 1:
//		Input: nestedList = [[1,1],2,[1,1]]
//		Output: [1,1,2,1,1]
//		Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,1,2,1,1].
//	Example 2:
//		Input: nestedList = [1,[4,[6]]]
//		Output: [1,4,6]
//		Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,4,6].
//	Constraints:
//		1 <= nestedList.length <= 500
//		The values of the integers in the nested list is in the range [-106, 106].

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor341(nestedList []*NestedInteger) *NestedIterator {
	l := len(nestedList)
	stack := make([]*NestedInteger, 0)
	for i := l - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	for len(stack) > 0 && !stack[len(stack)-1].IsInteger() {
		curNestedInteger := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nestedList := curNestedInteger.GetList()
		for i := len(nestedList) - 1; i >= 0; i-- {
			stack = append(stack, nestedList[i])
		}
	}
	return &NestedIterator{
		stack: stack,
	}
}

func (this *NestedIterator) Next() int {
	for this.HasNext() && !this.stack[len(this.stack)-1].IsInteger() {
		curNestedInteger := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		nestedList := curNestedInteger.GetList()
		for i := len(nestedList) - 1; i >= 0; i-- {
			this.stack = append(this.stack, nestedList[i])
		}
	}
	curNestedInteger := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return curNestedInteger.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 && !this.stack[len(this.stack)-1].IsInteger() {
		curNestedInteger := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		nestedList := curNestedInteger.GetList()
		for i := len(nestedList) - 1; i >= 0; i-- {
			this.stack = append(this.stack, nestedList[i])
		}
	}
	return len(this.stack) > 0
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct{}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return true }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 1 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }
