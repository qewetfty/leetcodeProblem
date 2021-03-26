package string

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// We are given two arrays A and B of words.  Each word is a string of lowercase letters.
// Now, say that word b is a subset of word a if every letter in b occurs in a,
// including multiplicity. For example, "wrr" is a subset of "warrior", but is
// not a subset of "world".
// Now say a word a from A is universal if for every b in B, b is a subset of a.
// Return a list of all universal words in A.  You can return the words in any order.
//	Example 1:
//		Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
//		Output: ["facebook","google","leetcode"]
//	Example 2:
//		Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
//		Output: ["apple","google","leetcode"]
//	Example 3:
//		Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
//		Output: ["facebook","google"]
//	Example 4:
//		Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
//		Output: ["google","leetcode"]
//	Example 5:
//		Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
//		Output: ["facebook","leetcode"]
//	Note:
//		1 <= A.length, B.length <= 10000
//		1 <= A[i].length, B[i].length <= 10
//		A[i] and B[i] consist only of lowercase letters.

// 使用数组不进行预处理，击败85.71%
func wordSubsets(A []string, B []string) []string {
	Bchar := make([]int, 26)
	for _, s := range B {
		Bcount, l := make([]int, 26), len(s)
		for i := 0; i < l; i++ {
			Bcount[s[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			Bchar[i] = utils.Max(Bchar[i], Bcount[i])
		}
	}
	result := make([]string, 0)

	for _, s := range A {
		Acount, l := make([]int, 26), len(s)
		for i := 0; i < l; i++ {
			Acount[s[i]-'a']++
		}
		find := true
		for i := 0; i < 26; i++ {
			if Acount[i] < Bchar[i] {
				find = false
				break
			}
		}
		if find {
			result = append(result, s)
		}
	}
	return result
}

// 带map预先处理，击败7.14%
func wordSubsets2(A []string, B []string) []string {
	Amap := make(map[string]map[byte]int)
	for _, s := range A {
		Amap[s] = make(map[byte]int)
		for i := 0; i < len(s); i++ {
			Amap[s][s[i]]++
		}
	}
	Bchar := make([]int, 26)
	for _, s := range B {
		Bcount := make([]int, 26)
		for i := 0; i < len(s); i++ {
			Bcount[s[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			Bchar[i] = utils.Max(Bchar[i], Bcount[i])
		}
	}

	result := make([]string, 0)
	// 比较A中每一个单词是否都包含B中的字母
	for s, smap := range Amap {
		find := true
		for i := 0; i < 26; i++ {
			if smap[byte(i)+'a'] < Bchar[i] {
				find = false
			}
		}
		if find {
			result = append(result, s)
		}
	}
	return result
}

func testProblem916() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"ec", "oc", "ceo"}))
}
