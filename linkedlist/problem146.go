package linkedlist

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

type dulNode struct {
	key   int
	value int
	prev  *dulNode
	next  *dulNode
}

func initDulNode(key, value int) *dulNode {
	return &dulNode{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	size     int
	capacity int
	head     *dulNode
	tail     *dulNode
	cache    map[int]*dulNode
}

func Constructor(capacity int) LRUCache {
	head := initDulNode(0, 0)
	tail := initDulNode(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*dulNode),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		curNode := this.cache[key]
		this.moveToHead(curNode)
		return curNode.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		curNode := this.cache[key]
		curNode.value = value
		this.moveToHead(curNode)
	} else {
		newNode := initDulNode(key, value)
		this.cache[key] = newNode
		this.addToHead(newNode)
		this.size++
		if this.size > this.capacity {
			removeNode := this.removeTail()
			this.size--
			delete(this.cache, removeNode.key)
		}
	}
}

func (this *LRUCache) addToHead(node *dulNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *dulNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *dulNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *dulNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
