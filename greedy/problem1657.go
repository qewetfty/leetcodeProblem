package greedy

import "fmt"

// Two strings are considered close if you can attain one from the other using the following operations:
//		Operation 1: Swap any two existing characters.
//					For example, abcde -> aecdb
//		Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
//					For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
//					You can use the operations on either string as many times as necessary.
//	Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.
//	Example 1:
//		Input: word1 = "abc", word2 = "bca"
//		Output: true
//		Explanation: You can attain word2 from word1 in 2 operations.
//					Apply Operation 1: "abc" -> "acb"
//					Apply Operation 1: "acb" -> "bca"
//	Example 2:
//		Input: word1 = "a", word2 = "aa"
//		Output: false
//		Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
//	Example 3:
//		Input: word1 = "cabbba", word2 = "abbccc"
//		Output: true
//		Explanation: You can attain word2 from word1 in 3 operations.
//					Apply Operation 1: "cabbba" -> "caabbb"
//					Apply Operation 2: "caabbb" -> "baaccc"
//					Apply Operation 2: "baaccc" -> "abbccc"
//	Example 4:
//		Input: word1 = "cabbba", word2 = "aabbss"
//		Output: false
//		Explanation: It is impossible to attain word2 from word1, or vice versa, in any amount of operations.
//	Constraints:
//		1 <= word1.length, word2.length <= 105
//		word1 and word2 contain only lowercase English letters.

// 判断 1.两个字符串是否包含相同的字符
//     2.所包含的字符中，包含的个数是否是相同的
func closeStrings(word1 string, word2 string) bool {
	l1, l2 := len(word1), len(word2)
	if l1 != l2 {
		return false
	}
	charMap := make([]int, 26)
	for i := 0; i < l1; i++ {
		charMap[word1[i]-'a']++
	}
	numberMap := make(map[int]int)
	for _, num := range charMap {
		if num > 0 {
			numberMap[num]++
		}
	}
	charMap2 := make([]int, 26)
	for i := 0; i < l2; i++ {
		// 出现了新的字母，直接false
		if charMap[word2[i]-'a'] == 0 {
			return false
		}
		charMap2[word2[i]-'a']++
	}
	for _, num := range charMap2 {
		if num == 0 {
			continue
		}
		if numberMap[num] == 0 {
			return false
		}
		numberMap[num]--
	}
	return true
}

func testProblem1657() {
	fmt.Println(closeStrings("abc", "bca"))
	fmt.Println(closeStrings("a", "aa"))
	fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("cabbba", "aabbss"))
}
