package hashTable

// Design and implement a data structure for a Least Frequently Used (LFU) cache.
//	Implement the LFUCache class:
//		LFUCache(int capacity) Initializes the object with the capacity of the data structure.
//		int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
//		void put(int key, int value) Update the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
// To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the smallest use counter is the least frequently used key.
// When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache is incremented either a get or put operation is called on it.
//	Example 1:
//		Input
//			["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
//			[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//		Output
//			[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//		Explanation
//			// cnt(x) = the use counter for key x
//			// cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
//			LFUCache lfu = new LFUCache(2);
//			lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
//			lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
//			lfu.get(1);      // return 1
//			                 // cache=[1,2], cnt(2)=1, cnt(1)=2
//			lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
//			                 // cache=[3,1], cnt(3)=1, cnt(1)=2
//			lfu.get(2);      // return -1 (not found)
//			lfu.get(3);      // return 3
//			                 // cache=[3,1], cnt(3)=2, cnt(1)=2
//			lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
//			                 // cache=[4,3], cnt(4)=1, cnt(3)=2
//			lfu.get(1);      // return -1 (not found)
//			lfu.get(3);      // return 3
//			                 // cache=[3,4], cnt(4)=1, cnt(3)=3
//			lfu.get(4);      // return 4
//			                 // cache=[3,4], cnt(4)=2, cnt(3)=3
//	Constraints:
//		0 <= capacity, key, value <= 104
//		At most 105 calls will be made to get and put.
//	Follow up: Could you do both operations in O(1) time complexity?

type LFUCache struct {
	capacity int
	size     int
	minFreq  int
	keyMap   map[int]*dulNode
	freqMap  map[int]*linkedList
}

func Constructor460(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		size:     0,
		minFreq:  0,
		keyMap:   make(map[int]*dulNode),
		freqMap:  make(map[int]*linkedList),
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.keyMap[key]; ok {
		this.IncreaseFreq(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if node, ok := this.keyMap[key]; ok {
		node.value = value
		this.IncreaseFreq(node)
	} else {
		if this.size >= this.capacity {
			node := this.freqMap[this.minFreq].removeLast()
			delete(this.keyMap, node.key)
			this.size--
		}
		newNode := createDulNode(key, value, 1)
		this.keyMap[key] = newNode
		if this.freqMap[1] == nil {
			this.freqMap[1] = newLinkedList()
		}
		this.freqMap[1].addFirst(newNode)
		this.minFreq = 1
		this.size++
	}
}

func (this *LFUCache) IncreaseFreq(node *dulNode) {
	curFreq := node.freq
	this.freqMap[curFreq].remove(node)
	if curFreq == this.minFreq && this.freqMap[curFreq].isEmpty() {
		this.minFreq++
		delete(this.freqMap, curFreq)
	}
	node.freq++
	if this.freqMap[node.freq] == nil {
		this.freqMap[node.freq] = newLinkedList()
	}
	this.freqMap[node.freq].addFirst(node)
}

type linkedList struct {
	head *dulNode
	tail *dulNode
}

func newLinkedList() *linkedList {
	head, tail := &dulNode{}, &dulNode{}
	head.next, tail.prev = tail, head
	return &linkedList{
		head: head,
		tail: tail,
	}
}

func (l *linkedList) addFirst(node *dulNode) {
	next := l.head.next
	node.prev, node.next = l.head, next
	l.head.next = node
	next.prev = node
}

func (l *linkedList) remove(node *dulNode) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil
}

func (l *linkedList) removeLast() *dulNode {
	if l.isEmpty() {
		return nil
	}
	removeNode := l.tail.prev
	l.remove(removeNode)
	return removeNode
}

func (l *linkedList) isEmpty() bool {
	return l.head.next == l.tail
}

type dulNode struct {
	key   int
	value int
	freq  int
	prev  *dulNode
	next  *dulNode
}

func createDulNode(key, value, freq int) *dulNode {
	return &dulNode{
		key:   key,
		value: value,
		freq:  freq,
	}
}
