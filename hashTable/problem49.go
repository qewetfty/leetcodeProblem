package hashTable

import (
	"fmt"
	"sort"
	"strconv"
)

// Given an array of strings, group anagrams together.
//	Example:
//		Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//		Output:
//			[
//  			["ate","eat","tea"],
//  			["nat","tan"],
//  			["bat"]
//			]
//	Note:
//		All inputs will be in lowercase.
//		The order of your output does not matter.

// sort string solution
func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	resMap := make(map[string][]string)
	for _, str := range strs {
		a := byteSlice(str)
		sort.Sort(a)
		keyStr := string(a)
		if v, exist := resMap[keyStr]; exist {
			v = append(v, str)
			resMap[keyStr] = v
		} else {
			value := make([]string, 0)
			value = append(value, str)
			resMap[keyStr] = value
		}
	}
	result := make([][]string, 0)
	for _, v := range resMap {
		result = append(result, v)
	}
	return result
}

type byteSlice []byte

func (b byteSlice) Len() int {
	return len(b)
}

func (b byteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b byteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// get char count solution
func groupAnagrams2(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	resMap := make(map[string][]string)
	for _, str := range strs {
		intSlice := make([]int, 26)
		for _, char := range str {
			intSlice[char-'a']++
		}
		intStr := convert2Str(intSlice)
		if v, exist := resMap[intStr]; exist {
			v = append(v, str)
			resMap[intStr] = v
		} else {
			value := make([]string, 0)
			value = append(value, str)
			resMap[intStr] = value
		}
	}
	result := make([][]string, 0)
	for _, v := range resMap {
		result = append(result, v)
	}
	return result
}

func convert2Str(intSlice []int) string {
	s := ""
	for _, n := range intSlice {
		s += strconv.Itoa(n)
	}
	return s
}

func testProblem49() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams2(a))
}
