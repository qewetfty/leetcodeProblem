package trie

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Design a special dictionary with some words that searchs the words in it by a prefix and a suffix.
// Implement the WordFilter class:
//	WordFilter(string[] words) Initializes the object with the words in the dictionary.
// 	f(string prefix, string suffix) Returns the index of the word in the
// 	dictionary, which has the prefix prefix and the suffix suffix. If there is
// 	more than one valid index, return the largest of them. If there is no such
// 	word in the dictionary, return -1.
//	Example 1:
//		Input
//			["WordFilter", "f"]
//			[[["apple"]], ["a", "e"]]
//		Output
//			[null, 0]
//		Explanation
//			WordFilter wordFilter = new WordFilter(["apple"]);
//			wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a" and suffix = 'e".
//	Constraints:
//		1 <= words.length <= 15000
//		1 <= words[i].length <= 10
//		1 <= prefix.length, suffix.length <= 10
//		words[i], prefix and suffix consist of lower-case English letters only.
//		At most 15000 calls will be made to the function f.

type WordFilter struct {
	t *trie745
}

func Constructor745(words []string) WordFilter {
	wordFilter := WordFilter{t: newTrie()}
	for i, word := range words {
		wordFilter.t.Insert("#"+word, i)
		l := len(word)
		for j := l - 1; j >= 0; j-- {
			wordFilter.t.Insert(word[j:]+"#"+word, i)
		}
	}
	return wordFilter
}

func (this *WordFilter) F(prefix string, suffix string) int {
	findWord := suffix + "#" + prefix
	find, index := this.t.Prefix(findWord)
	if !find {
		return -1
	}
	return index
}

type trie745 struct {
	node *trieNode745
}

func newTrie() *trie745 {
	return &trie745{
		node: newTrieNode745(),
	}
}

func (t *trie745) Insert(s string, index int) {
	n := t.node
	l := len(s)
	for i := 0; i < l; i++ {
		char := s[i]
		if !n.Contains(char) {
			n.Put(char, newTrieNode745(), index)
		}
		n, _ = n.Get(char)
		n.index = utils.Max(index, n.index)
	}
	n.isEnd = true
}

func (t *trie745) Prefix(prefix string) (bool, int) {
	n := t.node
	l := len(prefix)
	for i := 0; i < l; i++ {
		char := prefix[i]
		if !n.Contains(char) {
			return false, -1
		}
		n, _ = n.Get(char)
	}
	if n == nil {
		return false, -1
	}
	return true, n.index
}

func (t *trie745) Search(word string) (bool, int) {
	n := t.node
	l := len(word)
	for i := 0; i < l; i++ {
		char := word[i]
		if !n.Contains(char) {
			return false, -1
		}
		n, _ = n.Get(char)
	}
	find := n != nil && n.isEnd
	if find {
		return find, n.index
	}
	return find, -1
}

type trieNode745 struct {
	node  map[byte]*trieNode745
	isEnd bool
	index int
}

func newTrieNode745() *trieNode745 {
	return &trieNode745{node: make(map[byte]*trieNode745)}
}

func (n *trieNode745) Contains(char byte) bool {
	return n.node[char] != nil
}

func (n *trieNode745) Get(char byte) (*trieNode745, int) {
	return n.node[char], n.index
}

func (n *trieNode745) Put(char byte, node *trieNode745, index int) {
	node.index = index
	n.node[char] = node
}

func testProblem745() {
	w := Constructor745([]string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"})
	fmt.Println(w.F("a", "aa"))
	fmt.Println(w.F("ab", "abcaccbcaa"))
	fmt.Println(w.F("bccbacbcba", "a"))
}
