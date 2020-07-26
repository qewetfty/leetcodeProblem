package main

import "fmt"

// Given a 2D board and a list of words from the dictionary, find all words in the board.
// Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those
// horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.
//	Example:
//		Input:
//			board = [
//			  ['o','a','a','n'],
//			  ['e','t','a','e'],
//			  ['i','h','k','r'],
//			  ['i','f','l','v']
//			]
//			words = ["oath","pea","eat","rain"]
//		Output: ["eat","oath"]
//	Note:
//		All inputs are consist of lowercase letters a-z.
//		The values of words are distinct.

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	// build trie
	for _, word := range words {
		trie.Insert(word)
	}
	node := trie.root

	res := make([]string, 0)
	m, n := len(board), len(board[0])
	visitMap := make([][]bool, m)
	for i := 0; i < m; i++ {
		visitMap[i] = make([]bool, n)
	}

	// dfs search
	curWord := make([]byte, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if node.Contains(board[i][j]) {
				searchWords(i, j, board, &curWord, &visitMap, node.get(board[i][j]), &res)
			}
		}
	}
	return RemoveRepByMap(res)
}

// remove the duplicated words
func RemoveRepByMap(slc []string) []string {
	result := make([]string, 0)
	// save different keys
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		// if map length changed, we add a different word into map, then add it to the result
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func searchWords(m, n int, board [][]byte, curWord *[]byte, visitMap *[][]bool, node *TrieNode, res *[]string) {
	curByte := board[m][n]
	*curWord = append(*curWord, curByte)
	(*visitMap)[m][n] = true
	if node.isEnd {
		*res = append(*res, string(*curWord))
	}
	for i := 0; i < 4; i++ {
		newM, newN := m+dx[i], n+dy[i]
		if newM >= 0 && newN >= 0 && newM < len(board) && newN < len(board[0]) && !(*visitMap)[newM][newN] &&
			node.Contains(board[newM][newN]) {
			searchWords(m+dx[i], n+dy[i], board, curWord, visitMap, node.get(board[newM][newN]), res)
		}
	}
	*curWord = (*curWord)[:len(*curWord)-1]
	(*visitMap)[m][n] = false
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	links []*TrieNode
	isEnd bool
}

func NewTrieNode() *TrieNode {
	node := new(TrieNode)
	node.links = make([]*TrieNode, 26)
	return node
}

func (node *TrieNode) Contains(char byte) bool {
	return node.links[char-'a'] != nil
}

func (node *TrieNode) get(char byte) *TrieNode {
	return node.links[char-'a']
}

func (node *TrieNode) put(char byte, trieNode *TrieNode) {
	node.links[char-'a'] = trieNode
}

func (node *TrieNode) SetEnd() {
	node.isEnd = true
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := new(Trie)
	trie.root = NewTrieNode()
	return *trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for i := range word {
		if !node.Contains(word[i]) {
			node.put(word[i], NewTrieNode())
		}
		node = node.get(word[i])
	}
	node.SetEnd()
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for i := range word {
		if node.Contains(word[i]) {
			node = node.get(word[i])
		} else {
			node = nil
			break
		}
	}
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for i := range prefix {
		if node.Contains(prefix[i]) {
			node = node.get(prefix[i])
		} else {
			node = nil
			break
		}
	}
	return node != nil
}

func main() {
	fmt.Println(findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		[]string{"oath", "pea", "eat", "rain"}))
	fmt.Println(findWords([][]byte{{'a'}, {'a'}}, []string{"a"}))
	fmt.Println(findWords([][]byte{{'a', 'b'}, {'a', 'a'}}, []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}))
}
