package hashTable

//Given two strings s1 and s2, return true if s2 contains a permutation of s1,
//or false otherwise.
//
// In other words, return true if one of s1's permutations is the substring of
//s2.
//
//
// Example 1:
//
//
//Input: s1 = "ab", s2 = "eidbaooo"
//Output: true
//Explanation: s2 contains one permutation of s1 ("ba").
//
//
// Example 2:
//
//
//Input: s1 = "ab", s2 = "eidboaoo"
//Output: false
//
//
//
// Constraints:
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 and s2 consist of lowercase English letters.
//
// Related Topics Hash Table Two Pointers String Sliding Window 👍 4382 👎 126

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	// l1严格小于l2
	if l1 > l2 {
		return false
	}
	s1Map, curS2Map := make([]int, 26), make([]int, 26)
	for i := 0; i < l1; i++ {
		s1Map[s1[i]-'a']++
		curS2Map[s2[i]-'a']++
	}
	// 先比较一次
	if compareString(s1Map, curS2Map) {
		return true
	}
	for i := l1; i < l2; i++ {
		// 更新s2的当前字母集
		curS2Map[s2[i-l1]-'a']--
		curS2Map[s2[i]-'a']++
		if compareString(s1Map, curS2Map) {
			return true
		}
	}
	return false
}

func compareString(m1, m2 []int) bool {
	find := true
	for i := 0; i < 26; i++ {
		if m1[i] != m2[i] {
			find = false
			break
		}
	}
	return find
}

//leetcode submit region end(Prohibit modification and deletion)
