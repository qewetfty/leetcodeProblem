package linkedlist

// Given an Iterator class interface with methods: next() and hasNext(), design
// and implement a PeekingIterator that support the peek() operation -- it
// essentially peek() at the element that will be returned by the next call to
// next().
//	Example:
//		Assume that the iterator is initialized to the beginning of the list: [1,2,3].
//		Call next() gets you 1, the first element in the list.
//		Now you call peek() and it returns 2, the next element. Calling next() after that still return 2.
//		You call next() the final time and it returns 3, the last element.
//		Calling hasNext() after that should return false.
//	Follow up: How would you extend your design to be generic and work with all types, not just integer?

type PeekingIterator struct {
	iter      *Iterator
	nextExist bool
	cur       int
}

func Constructor284(iter *Iterator) *PeekingIterator {
	var curNumber int
	exist := iter.hasNext()
	if iter.hasNext() {
		curNumber = iter.next()
	} else {
		curNumber = -1
	}
	return &PeekingIterator{iter: iter, nextExist: exist, cur: curNumber}
}

func (this *PeekingIterator) hasNext() bool {
	return this.nextExist
}

func (this *PeekingIterator) next() int {
	result := this.cur
	this.nextExist = this.iter.hasNext()
	var curNumber int
	if this.iter.hasNext() {
		curNumber = this.iter.next()
	} else {
		curNumber = -1
	}
	this.cur = curNumber
	return result
}

func (this *PeekingIterator) peek() int {
	return this.cur
}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	return true
	// Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	return 1
	// Returns the next element in the iteration.
}
