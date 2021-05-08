package main

import (
	"fmt"
	"strconv"
)

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
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*dulNode),
		head:     initDulNode(0, 0),
		tail:     initDulNode(0, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	return 1
}

func (this *LRUCache) Put(key int, value int) {

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

func superpalindromesInRange2(left string, right string) int {
	l, _ := strconv.ParseInt(left, 10, 64)
	r, _ := strconv.ParseInt(right, 10, 64)
	result := 0
	m := 10000
	// 计算奇数
	for i := 1; i < m; i++ {
		s := []byte(strconv.Itoa(i))
		for j := len(s) - 2; j >= 0; j-- {
			s = append(s, s[j])
		}
		n, _ := strconv.ParseInt(string(s), 10, 64)
		n *= n
		if n > r {
			break
		}
		if n >= l && isPalindrome(n) {
			result++
		}
	}
	// 计算偶数
	for i := 1; i < m; i++ {
		s := []byte(strconv.Itoa(i))
		for j := len(s) - 1; j >= 0; j-- {
			s = append(s, s[j])
		}
		n, _ := strconv.ParseInt(string(s), 10, 64)
		n *= n
		if n > r {
			break
		}
		if n >= l && isPalindrome(n) {
			result++
		}
	}
	return result
}

func isPalindrome(n int64) bool {
	return n == reverseLong(n)
}

func reverseLong(n int64) int64 {
	var result int64
	result = 0
	for n > 0 {
		result = 10*result + n%10
		n = n / 10
	}
	return result
}

func main() {
	fmt.Println(superpalindromesInRange2("1000000000000000", "100000000000000000"))
}
