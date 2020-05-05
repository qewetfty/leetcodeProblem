package dfs_bfs

import (
	"container/list"
	"fmt"
	"strings"
)

// A gene string can be represented by an 8-character long string, with choices from "A", "C", "G", "T".
// Suppose we need to investigate about a mutation (mutation from "start" to "end"), where ONE mutation is defined as
// ONE single character changed in the gene string.
// For example, "AACCGGTT" -> "AACCGGTA" is 1 mutation.
// Also, there is a given gene "bank", which records all the valid gene mutations. A gene must be in the bank to make it a valid gene string.
// Now, given 3 things - start, end, bank, your task is to determine what is the minimum number of mutations needed to
// mutate from "start" to "end". If there is no such a mutation, return -1.
//	Note:
//		Starting point is assumed to be valid, so it might not be included in the bank.
//		If multiple mutations are needed, all mutations during in the sequence must be valid.
//		You may assume start and end string is not the same.
//	Example 1:
//		start: "AACCGGTT"
//		end:   "AACCGGTA"
//		bank: ["AACCGGTA"]
//		return: 1
//	Example 2:
//		start: "AACCGGTT"
//		end:   "AAACGGTA"
//		bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
//		return: 2
//	Example 3:
//		start: "AAAAACCC"
//		end:   "AACCCCCC"
//		bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
//		return: 3

// BFS method
func minMutation(start string, end string, bank []string) int {
	if strings.EqualFold(start, end) {
		return 0
	}
	geneMap := make(map[string]int)
	for _, gene := range bank {
		geneMap[gene] = 0
	}
	if _, ok := geneMap[end]; !ok {
		return -1
	}
	deque := new(list.List)
	deque.PushBack(start)
	base := []string{"A", "T", "C", "G"}
	count := 0
	for deque.Len() > 0 {
		length := deque.Len()
		for i := 0; i < length; i++ {
			currDna := deque.Remove(deque.Front()).(string)
			if strings.EqualFold(currDna, end) {
				return count
			}
			for j := 0; j < len(currDna); j++ {
				for _, b := range base {
					newDna := currDna[:j] + b + currDna[j+1:]
					if _, ok := geneMap[newDna]; ok {
						deque.PushBack(newDna)
						delete(geneMap, newDna)
					}
				}
			}
		}
		count++
	}
	return -1
}

func testProblem433() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
}
