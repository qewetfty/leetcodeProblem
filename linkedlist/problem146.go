package linkedlist

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"strconv"
)

// Design and implement a data structure for Least Recently Used (LRU) cache. It should support
// the following operations: get and put.
//
// get(key) - Get the value (will always be positive) of the key if the key exists in the cache,
// otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present.
// When the cache reached its capacity, it should invalidate the least recently used item before
// inserting a new item.
//
// The cache is initialized with a positive capacity.

type LRUCache struct {
	capacity int
	count    int
	dataMap  map[int]*data.DulNode
	head     *data.DulNode
	tail     *data.DulNode
}

func Constructor(capacity int) LRUCache {
	cache := new(LRUCache)
	cache.capacity = capacity
	cache.count = 0
	cache.dataMap = make(map[int]*data.DulNode, 0)
	cache.head = new(data.DulNode)
	cache.tail = cache.head
	return *cache
}

func (this *LRUCache) Insert(node *data.DulNode) {
	if node == nil {
		return
	}
	next := this.head.Next
	this.head.Next = node
	node.Prev = this.head
	node.Next = next
	if this.tail == this.head {
		this.tail = node
	} else {
		next.Prev = node
	}
	this.count++
}

func (this *LRUCache) Delete(node *data.DulNode) {
	prev := node.Prev
	next := node.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		this.tail = prev
	}
	this.count--
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.dataMap[key]; ok {
		this.Delete(node)
		this.Insert(node)
		return node.Value
	} else {
		return -1
	}
}

func (this *LRUCache) ToString() string {
	resultStr := ""
	current := this.head.Next
	for current != nil {
		resultStr += strconv.Itoa(current.Value)
		if current != this.tail {
			resultStr += "->"
		}
		current = current.Next
	}
	return resultStr
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.dataMap[key]; ok {
		node.Value = value
		this.Delete(node)
		this.Insert(node)
		return
	}
	newNode := new(data.DulNode)
	newNode.Key = key
	newNode.Value = value
	if this.count < this.capacity {
		this.Insert(newNode)
		this.dataMap[key] = newNode
		return
	}
	delNode := this.tail
	this.Delete(delNode)
	delete(this.dataMap, delNode.Key)
	this.dataMap[key] = newNode
	this.Insert(newNode)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func testProblem146() {
	obj := Constructor(2)
	fmt.Println(obj.Get(2))
	obj.Put(2, 6)
	fmt.Println(obj.ToString())

	fmt.Println(obj.Get(1))

	fmt.Println(obj.ToString())
	obj.Put(1, 5)
	fmt.Println(obj.ToString())

	obj.Put(1, 2)

	fmt.Println(obj.Get(1))

	fmt.Println(obj.ToString())

	fmt.Println(obj.Get(2))

	fmt.Println(obj.ToString())
}
