package dfs_bfs

import (
	"container/list"
	"fmt"
	"strings"
)

// Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation
// sequence from beginWord to endWord, such that:
// Only one letter can be changed at a time.
// Each transformed word must exist in the word list.
//	Note:
//		Return 0 if there is no such transformation sequence.
//		All words have the same length.
//		All words contain only lowercase alphabetic characters.
//		You may assume no duplicates in the word list.
//		You may assume beginWord and endWord are non-empty and are not the same.
//	Example 1:
//		Input:
//			beginWord = "hit",
//			endWord = "cog",
//			wordList = ["hot","dot","dog","lot","log","cog"]
//		Output: 5
//	Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//	return its length 5.

//	Example 2:
//		Input:
//			beginWord = "hit"
//			endWord = "cog"
//			wordList = ["hot","dot","dog","lot","log"]
//		Output: 0
//	Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

// BFS method
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// judge whether a word is in word list and build search map
	searchMap := make(map[string][]string, 0)
	l := len(beginWord)
	contains := false
	for _, word := range wordList {
		if strings.EqualFold(endWord, word) {
			contains = true
		}
		for i := 0; i < l; i++ {
			pattern := word[:i] + "*" + word[i+1:]
			if _, ok := searchMap[pattern]; ok {
				searchList := searchMap[pattern]
				searchList = append(searchList, word)
				searchMap[pattern] = searchList
			} else {
				searchMap[pattern] = []string{word}
			}
		}
	}
	if !contains {
		return 0
	}

	// BFS search every possible state
	deque := new(list.List)
	deque.PushBack(beginWord)
	visitedMap := make(map[string]bool)
	visitedMap[beginWord] = true
	count := 1
	for deque.Len() > 0 {
		length := deque.Len()
		for i := 0; i < length; i++ {
			currWord := deque.Remove(deque.Front()).(string)
			for j := 0; j < l; j++ {
				pattern := currWord[:j] + "*" + currWord[j+1:]
				if _, ok := searchMap[pattern]; ok {
					words := searchMap[pattern]
					for _, word := range words {
						if strings.EqualFold(endWord, currWord) {
							return count
						}
						if _, ok := visitedMap[word]; !ok {
							visitedMap[word] = true
							deque.PushBack(word)
						}
					}
				}
			}
		}
		count++
	}
	return 0
}

func testProblem127() {
	a := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength("hit", "cog", a))
	fmt.Println(ladderLength("a", "c", []string{"a", "b", "c"}))
	fmt.Println(ladderLength("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
}
