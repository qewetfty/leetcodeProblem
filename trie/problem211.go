package trie

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Design a data structure that supports the following two operations:
//	void addWord(word)
//	bool search(word)
//	search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.
//	Example:
//		addWord("bad")
//		addWord("dad")
//		addWord("mad")
//		search("pad") -> false
//		search("bad") -> true
//		search(".ad") -> true
//		search("b..") -> true
//	Note:
//		You may assume that all words are consist of lowercase letters a-z.

type WordDictionary struct {
	trie data.Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		trie: data.NewTrie(),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return search(word, 0, this.trie.Root)
}

func search(word string, index int, node *data.TrieNode) bool {
	if index == len(word) {
		if node.IsEnd {
			return true
		} else {
			return false
		}
	}
	if word[index] == '.' {
		find := false
		for char := byte('a'); char <= byte('z'); char++ {
			if node.Contains(char) {
				find = search(word, index+1, node.Get(char))
			}
			if find {
				return find
			}
		}
		return false
	} else {
		if node.Contains(word[index]) {
			return search(word, index+1, node.Get(word[index]))
		} else {
			return false
		}
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func testProblem211() {
	obj := Constructor()
	obj.AddWord("at")
	obj.AddWord("and")
	obj.AddWord("an")
	obj.AddWord("add")
	fmt.Println(obj.Search("a"))
	fmt.Println(obj.Search(".at"))
	obj.AddWord("bat")
	fmt.Println(obj.Search(".at"))
	fmt.Println(obj.Search("an."))
	fmt.Println(obj.Search("a.d."))
	fmt.Println(obj.Search("a.d."))
	fmt.Println(obj.Search("."))
}
