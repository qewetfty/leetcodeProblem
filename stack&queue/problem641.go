package stack_queue

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Design your implementation of the circular double-ended queue (deque).
// Your implementation should support following operations:
//	MyCircularDeque(k): Constructor, set the size of the deque to be k.
//	insertFront(): Adds an item at the front of Deque. Return true if the operation is successful.
//	insertLast(): Adds an item at the rear of Deque. Return true if the operation is successful.
//	deleteFront(): Deletes an item from the front of Deque. Return true if the operation is successful.
//	deleteLast(): Deletes an item from the rear of Deque. Return true if the operation is successful.
//	getFront(): Gets the front item from the Deque. If the deque is empty, return -1.
//	getRear(): Gets the last item from Deque. If the deque is empty, return -1.
//	isEmpty(): Checks whether Deque is empty or not.
//	isFull(): Checks whether Deque is full or not.

type MyCircularDeque struct {
	capacity int
	count    int
	head     *data.DulNode
	tail     *data.DulNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func MyCircularDequeConstructor(k int) MyCircularDeque {
	m := new(MyCircularDeque)
	m.capacity = k
	m.head = new(data.DulNode)
	m.tail = new(data.DulNode)
	m.head.Next = m.tail
	m.tail.Prev = m.head
	return *m
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.count == this.capacity {
		return false
	}
	node := new(data.DulNode)
	node.Value = value
	next := this.head.Next
	this.head.Next = node
	next.Prev = node
	node.Prev = this.head
	node.Next = next
	this.count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.count == this.capacity {
		return false
	}
	node := new(data.DulNode)
	node.Value = value
	prev := this.tail.Prev
	this.tail.Prev = node
	prev.Next = node
	node.Prev = prev
	node.Next = this.tail
	this.count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.count == 0 {
		return false
	}
	del := this.head.Next
	this.head.Next = del.Next
	del.Next.Prev = this.head
	del.Next = nil
	del.Prev = nil
	del = nil
	this.count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.count == 0 {
		return false
	}
	del := this.tail.Prev
	this.tail.Prev = del.Prev
	del.Prev.Next = this.tail
	del.Next = nil
	del.Prev = nil
	del = nil
	this.count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.count == 0 {
		return -1
	}
	return this.head.Next.Value
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.count == 0 {
		return -1
	}
	return this.tail.Prev.Value
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := MyCircularDequeConstructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func testProblem641() {
	m := MyCircularDequeConstructor(3)
	fmt.Println(m.InsertFront(9))
	fmt.Println(m.GetRear())
	fmt.Println(m.InsertFront(9))
	fmt.Println(m.GetRear())
	fmt.Println(m.InsertLast(5))
	fmt.Println(m.GetFront())
	fmt.Println(m.GetRear())
	fmt.Println(m.InsertLast(8))
	fmt.Println(m.DeleteLast())
	fmt.Println(m.GetFront())
}
