package dfs_bfs

import (
	"container/list"
	"fmt"
	"strings"
)

// Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s)
// from beginWord to endWord, such that:
//		Only one letter can be changed at a time
//		Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
//	Note:
//		Return an empty list if there is no such transformation sequence.
//		All words have the same length.
//		All words contain only lowercase alphabetic characters.
//		You may assume no duplicates in the word list.
//		You may assume beginWord and endWord are non-empty and are not the same.
//	Example 1:
//		Input:
//			beginWord = "hit",
//			endWord = "cog",
//			wordList = ["hot","dot","dog","lot","log","cog"]
//
//		Output:
//			[
//			  ["hit","hot","dot","dog","cog"],
//			  ["hit","hot","lot","log","cog"]
//			]
//	Example 2:
//		Input:
//			beginWord = "hit"
//			endWord = "cog"
//			wordList = ["hot","dot","dog","lot","log"]
//		Output: []
//	Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
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
		return res
	}

	// BFS search every possible state
	deque := new(list.List)
	deque.PushBack([]string{beginWord})
	visitedMap := make(map[string]bool)
	visitedMap[beginWord] = true
	find := false
	for deque.Len() > 0 && !find {
		length := deque.Len()
		subVisitList := make([]string, 0)
		for i := 0; i < length; i++ {
			currWordList := deque.Remove(deque.Front()).([]string)
			currWord := currWordList[len(currWordList)-1]
			for j := 0; j < l; j++ {
				pattern := currWord[:j] + "*" + currWord[j+1:]
				if _, ok := searchMap[pattern]; !ok {
					continue
				}
				words := searchMap[pattern]
				for _, word := range words {
					// copy slice of visited path
					if strings.EqualFold(word, endWord) {
						newVisitedPath := make([]string, 0)
						for _, s := range currWordList {
							newVisitedPath = append(newVisitedPath, s)
						}
						newVisitedPath = append(newVisitedPath, word)
						res = append(res, newVisitedPath)
						find = true
						continue
					}
					if _, visited := visitedMap[word]; !visited {
						newVisitedPath := make([]string, 0)
						for _, s := range currWordList {
							newVisitedPath = append(newVisitedPath, s)
						}
						newVisitedPath = append(newVisitedPath, word)
						subVisitList = append(subVisitList, word)
						deque.PushBack(newVisitedPath)
					}
				}
			}
		}
		for _, s := range subVisitList {
			visitedMap[s] = true
		}
	}
	return res
}

func testProblem126() {
	fmt.Println(findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
	a := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(findLadders("hit", "cog", a))
	fmt.Println(findLadders("a", "c", []string{"a", "b", "c"}))
}
