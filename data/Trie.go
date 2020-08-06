package data

type Trie struct {
	Root *TrieNode
}

func NewTrie() Trie {
	return Trie{
		Root: NewTrieNode(),
	}
}

func (trie *Trie) Insert(word string) {
	node := trie.Root
	for i := 0; i < len(word); i++ {
		if !node.Contains(word[i]) {
			node.Put(word[i], NewTrieNode())
		}
		node = node.Get(word[i])
	}
	node.IsEnd = true
}

func (trie *Trie) Search(word string) bool {
	node := trie.Root
	for i := 0; i < len(word); i++ {
		if node.Contains(word[i]) {
			node = node.Get(word[i])
		} else {
			node = nil
			break
		}
	}
	return node != nil && node.IsEnd
}

func (trie *Trie) StartWith(prefix string) bool {
	node := trie.Root
	for i := 0; i < len(prefix); i++ {
		if node.Contains(prefix[i]) {
			node = node.Get(prefix[i])
		} else {
			node = nil
			break
		}
	}
	return node != nil
}

type TrieNode struct {
	List  []*TrieNode
	IsEnd bool
}

func NewTrieNode() *TrieNode {
	node := new(TrieNode)
	node.List = make([]*TrieNode, 26)
	return node
}

func (node *TrieNode) Contains(char byte) bool {
	return node.List[char-'a'] != nil
}

func (node *TrieNode) Get(char byte) *TrieNode {
	return node.List[char-'a']
}

func (node *TrieNode) Put(char byte, trieNode *TrieNode) {
	node.List[char-'a'] = trieNode
}
