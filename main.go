package main

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// 超时
type WordFilter struct {
	prefixTrie *trie
	suffixTrie *trie
}

func Constructor(words []string) WordFilter {
	w := WordFilter{
		prefixTrie: NewTrie(),
		suffixTrie: NewTrie(),
	}
	for i, word := range words {
		w.prefixTrie.Insert(word, i)
		reverseWord := ReverseString(word)
		w.suffixTrie.Insert(reverseWord, i)
	}
	return w
}

func (this *WordFilter) F(prefix string, suffix string) int {
	findPrefix, prefixIndex := this.prefixTrie.StartWith(prefix)
	if !findPrefix {
		return -1
	}
	reverseSuffix := ReverseString(suffix)
	findSuffix, suffixIndex := this.suffixTrie.StartWith(reverseSuffix)
	if !findSuffix {
		return -1
	}
	// 遍历其中一个map，如果所有的都不在另一个map，直接返回-1
	result := -1
	for i := range prefixIndex {
		if suffixIndex[i] {
			result = utils.Max(result, i)
		}
	}
	return result
}

func ReverseString(s string) string {
	r := []rune(s)
	for from, to := 0, len(r)-1; from < to; from, to = from+1, to-1 {
		r[from], r[to] = r[to], r[from]
	}
	return string(r)
}

type trie struct {
	t *trieNode
}

func NewTrie() *trie {
	return &trie{
		t: MakeNewTrieNode(),
	}
}

func (t *trie) Insert(s string, index int) {
	node := t.t
	l := len(s)
	for i := 0; i < l; i++ {
		char := s[i]
		if !node.Contains(char) {
			node.Put(char, MakeNewTrieNode(), index)
		}
		node.PutIndex(index)
		node = node.Get(char)
	}
	node.isEnd = true
	node.PutIndex(index)
}

func (t *trie) StartWith(prefix string) (bool, map[int]bool) {
	node := t.t
	l := len(prefix)
	for i := 0; i < l; i++ {
		char := prefix[i]
		if node.Contains(char) {
			node = node.Get(char)
		} else {
			node = nil
			break
		}
	}
	if node == nil {
		return false, map[int]bool{}
	}
	return true, node.indexes
}

func MakeNewTrieNode() *trieNode {
	return &trieNode{
		charMap: make(map[byte]*trieNode),
		indexes: make(map[int]bool),
	}
}

type trieNode struct {
	charMap map[byte]*trieNode
	isEnd   bool
	indexes map[int]bool
}

func (t *trieNode) Contains(char byte) bool {
	return t.charMap[char] != nil
}

func (t *trieNode) Get(char byte) *trieNode {
	return t.charMap[char]
}

func (t *trieNode) Put(char byte, node *trieNode, index int) {
	t.charMap[char] = node
	t.indexes[index] = true
}

func (t *trieNode) PutIndex(index int) {
	t.indexes[index] = true
}

func test() {
	w := Constructor([]string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"})
	fmt.Println(w.F("bccbacbcba", "a"))
	fmt.Println(w.F("ab", "abcaccbcaa"))
}
