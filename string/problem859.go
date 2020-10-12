package string

import "fmt"

// Given two strings A and B of lowercase letters, return true if you can swap
// two letters in A so the result is equal to B, otherwise, return false.
// Swapping letters is defined as taking two indices i and j (0-indexed) such
// that i != j and swapping the characters at A[i] and A[j]. For example,
// swapping at indices 0 and 2 in "abcd" results in "cbad".
//	Example 1:
//		Input: A = "ab", B = "ba"
//		Output: true
//		Explanation: You can swap A[0] = 'a' and A[1] = 'b' to get "ba", which is equal to B.
//	Example 2:
//		Input: A = "ab", B = "ab"
//		Output: false
//		Explanation: The only letters you can swap are A[0] = 'a' and A[1] = 'b', which results in "ba" != B.
//	Example 3:
//		Input: A = "aa", B = "aa"
//		Output: true
//		Explanation: You can swap A[0] = 'a' and A[1] = 'a' to get "aa", which is equal to B.
//	Example 4:
//		Input: A = "aaaaaaabc", B = "aaaaaaacb"
//		Output: true
//	Example 5:
//		Input: A = "", B = "aa"
//		Output: false
//	Constraints:
//		0 <= A.length <= 20000
//		0 <= B.length <= 20000
//		A and B consist of lowercase letters.

func buddyStrings(A string, B string) bool {
	l1, l2 := len(A), len(B)
	// 长度不相同，直接返回false
	if l1 != l2 {
		return false
	}
	diff, charMap := make([]int, 0), make(map[byte]int)
	// 记录不相同的位置，和每个字符出现的次数
	for i := 0; i < l1; i++ {
		if A[i] != B[i] {
			diff = append(diff, i)
		}
		charMap[A[i]]++
	}
	// 如果字符串不相同位置的长度不是2或者0，直接返回false
	if len(diff) != 0 && len(diff) != 2 {
		return false
	}
	// 如果diff长度为2，则判断两个不相同位置的字符是否可交换
	if len(diff) == 2 {
		if A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]] {
			return true
		} else {
			return false
		}
	}
	// 如果diff长度为0，则判断是否有相同的字符，有则返回true，无则返回false
	for _, val := range charMap {
		if val > 1 {
			return true
		}
	}
	return false
}

func testProblem859() {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Println(buddyStrings("", "aa"))
}
