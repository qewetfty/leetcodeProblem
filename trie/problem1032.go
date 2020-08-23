package trie

import "fmt"

// Implement the StreamChecker class as follows:
// StreamChecker(words): Constructor1032, init the data structure with the given words.
// query(letter): returns true if and only if for some k >= 1, the last k characters queried (in order from oldest to
// newest, including this letter just queried) spell one of the words in the given list.
//	Example:
//		StreamChecker streamChecker = new StreamChecker(["cd","f","kl"]); // init the dictionary.
//		streamChecker.query('a');          // return false
//		streamChecker.query('b');          // return false
//		streamChecker.query('c');          // return false
//		streamChecker.query('d');          // return true, because 'cd' is in the wordlist
//		streamChecker.query('e');          // return false
//		streamChecker.query('f');          // return true, because 'f' is in the wordlist
//		streamChecker.query('g');          // return false
//		streamChecker.query('h');          // return false
//		streamChecker.query('i');          // return false
//		streamChecker.query('j');          // return false
//		streamChecker.query('k');          // return false
//		streamChecker.query('l');          // return true, because 'kl' is in the wordlist
//	Note:
//		1 <= words.length <= 2000
//		1 <= words[i].length <= 2000
//		Words will only consist of lowercase English letters.
//		Queries will only consist of lowercase English letters.
//		The number of queries is at most 40000.

type StreamChecker struct {
	Root     *trienode
	charList []byte
}

func Constructor1032(words []string) StreamChecker {
	checker := new(StreamChecker)
	checker.Root = newTrieNode1032()
	// 构建后缀树
	for _, word := range words {
		node, curCharList := checker.Root, []byte(word)
		for i := len(word) - 1; i >= 0; i-- {
			if !node.Contains(curCharList[i]) {
				node.Put(curCharList[i], newTrieNode1032())
			}
			node = node.Get(curCharList[i])
		}
		node.isEnd = true
	}
	return *checker
}

func (this *StreamChecker) Query(letter byte) bool {
	this.charList = append(this.charList, letter)
	node := this.Root
	for i := len(this.charList) - 1; i >= 0; i-- {
		if !node.Contains(this.charList[i]) {
			return false
		}
		node = node.Get(this.charList[i])
		if node.isEnd {
			return true
		}
	}
	return false
}

type trienode struct {
	list  []*trienode
	isEnd bool
}

func newTrieNode1032() *trienode {
	node := new(trienode)
	node.list = make([]*trienode, 26)
	return node
}

func (t *trienode) Contains(char byte) bool {
	return t.list[char-'a'] != nil
}

func (t *trienode) Get(char byte) *trienode {
	return t.list[char-'a']
}

func (t *trienode) Put(char byte, node *trienode) {
	t.list[char-'a'] = node
}

func testProblem1032() {
	checker := Constructor1032([]string{"abaa", "abaab", "aabbb", "bab", "ab"})
	fmt.Println(checker.Query('a'))
	fmt.Println(checker.Query('a'))
	fmt.Println(checker.Query('b'))
	fmt.Println(checker.Query('b'))
	fmt.Println(checker.Query('b'))
}
