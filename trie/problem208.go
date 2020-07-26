package trie

// Implement a trie with insert, search, and startsWith methods.
//	Example:
//		Trie trie = new Trie();
//		trie.insert("apple");
//		trie.search("apple");   // returns true
//		trie.search("app");     // returns false
//		trie.startsWith("app"); // returns true
//		trie.insert("app");
//		trie.search("app");     // returns true
//	Note:
//		You may assume that all inputs are consist of lowercase letters a-z.
//		All inputs are guaranteed to be non-empty strings.

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

func (node *TrieNode) Contains(char int32) bool {
	return node.links[char-'a'] != nil
}

func (node *TrieNode) get(char int32) *TrieNode {
	return node.links[char-'a']
}

func (node *TrieNode) put(char int32, trieNode *TrieNode) {
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
	for _, c := range word {
		if !node.Contains(c) {
			node.put(c, NewTrieNode())
		}
		node = node.get(c)
	}
	node.SetEnd()
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		if node.Contains(c) {
			node = node.get(c)
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
	for _, c := range prefix {
		if node.Contains(c) {
			node = node.get(c)
		} else {
			node = nil
			break
		}
	}
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func testProblem208() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Search("apple")
}
