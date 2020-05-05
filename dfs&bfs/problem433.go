package dfs_bfs

import (
	"container/list"
	"fmt"
	"math"
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

// DFS method
var base = []string{"A", "T", "C", "G"}

func minMutation2(start string, end string, bank []string) int {
	if strings.EqualFold(start, end) {
		return 0
	}
	geneMap := make(map[string]bool, len(bank))
	for _, gene := range bank {
		geneMap[gene] = false
	}
	if _, ok := geneMap[end]; !ok {
		return -1
	}
	var visitedMap = make(map[string]bool)
	var minCount = math.MaxInt32
	dfs433(start, end, 0, &minCount, &geneMap, &visitedMap)
	if minCount == math.MaxInt32 {
		minCount = -1
	}
	return minCount
}

func dfs433(start, end string, count int, minCount *int, geneMap, visitedMap *map[string]bool) {
	if start == end {
		*minCount = int(math.Min(float64(*minCount), float64(count)))
		return
	}
	if _, ok := (*visitedMap)[start]; ok {
		return
	}
	(*visitedMap)[start] = true
	for i := 0; i < len(start); i++ {
		for _, nu := range base {
			newDna := start[:i] + nu + start[i+1:]
			if _, ok := (*geneMap)[newDna]; ok {
				dfs433(newDna, end, count+1, minCount, geneMap, visitedMap)
			}
		}
	}
}

func testProblem433() {
	fmt.Println(minMutation2("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation2("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation2("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
}
