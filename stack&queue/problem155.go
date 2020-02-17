package stack_queue

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Design a stack&queue that supports push, pop, top, and retrieving the minimum element in constant time.
//	push(x) -- Push element x onto stack&queue.
//	pop() -- Removes the element on top of the stack&queue.
//	top() -- Get the top element.
//	getMin() -- Retrieve the minimum element in the stack&queue.
//
//	Example:
//		MinStack minStack = new MinStack();
//		minStack.push(-2);
//		minStack.push(0);
//		minStack.push(-3);
//		minStack.getMin();   --> Returns -3.
//		minStack.pop();
//		minStack.top();      --> Returns 0.
//		minStack.getMin();   --> Returns -2.

type MinStack struct {
	Data *data.Stack
	Min  *data.Stack
}

func MinStackConstructor() MinStack {
	minStack := new(MinStack)
	min := new(data.Stack)
	data := new(data.Stack)
	minStack.Data = data
	minStack.Min = min
	return *minStack
}

func (this *MinStack) Push(x int) {
	this.Data.Push(x)
	if this.Min.IsEmpty() || this.Min.Top().(int) >= x {
		this.Min.Push(x)
	}
}

func (this *MinStack) Pop() {
	val, _ := this.Data.Pop()
	if this.Min.Top().(int) == val {
		this.Min.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.Data.Top().(int)
}

func (this *MinStack) GetMin() int {
	return this.Min.Top().(int)
}

func testProblem155() {
	minStack := MinStackConstructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
