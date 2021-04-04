package stack_queue

// Design your implementation of the circular queue. The circular queue is a
// linear data structure in which the operations are performed based on FIFO
// (First In First Out) principle and the last position is connected back to the
// first position to make a circle. It is also called "Ring Buffer".
// One of the benefits of the circular queue is that we can make use of the
// spaces in front of the queue. In a normal queue, once the queue becomes full,
// we cannot insert the next element even if there is a space in front of the
// queue. But using the circular queue, we can use the space to store new values.
// Implementation the MyCircularQueue class:
//	MyCircularQueue(k) Initializes the object with the size of the queue to be k.
//	int Front() Gets the front item from the queue. If the queue is empty, return -1.
//	int Rear() Gets the last item from the queue. If the queue is empty, return -1.
//	boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
//	boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
//	boolean isEmpty() Checks whether the circular queue is empty or not.
//	boolean isFull() Checks whether the circular queue is full or not.
//	Example 1:
//		Input
//			["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
//			[[3], [1], [2], [3], [4], [], [], [], [4], []]
//		Output:
//			[null, true, true, true, false, 3, true, true, true, 4]
//		Explanation:
//			MyCircularQueue myCircularQueue = new MyCircularQueue(3);
//			myCircularQueue.enQueue(1); // return True
//			myCircularQueue.enQueue(2); // return True
//			myCircularQueue.enQueue(3); // return True
//			myCircularQueue.enQueue(4); // return False
//			myCircularQueue.Rear();     // return 3
//			myCircularQueue.isFull();   // return True
//			myCircularQueue.deQueue();  // return True
//			myCircularQueue.enQueue(4); // return True
//			myCircularQueue.Rear();     // return 4

type MyCircularQueue struct {
	queue []int
	start int
	end   int
	full  bool
}

func Constructor622(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		start: 0,
		end:   0,
		full:  k == 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	next := this.end + 1
	if next == len(this.queue) {
		next = 0
	}
	if next == this.start {
		this.full = true
	}
	this.end, this.queue[this.end] = next, value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.full = false
	prev := this.start + 1
	if prev == len(this.queue) {
		prev = 0
	}
	this.start = prev
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.start]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	prev := this.end - 1
	if prev < 0 {
		prev = len(this.queue) - 1
	}
	return this.queue[prev]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.start == this.end && !this.full
}

func (this *MyCircularQueue) IsFull() bool {
	return this.start == this.end && this.full
}
