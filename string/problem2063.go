package string

import "fmt"

// Given a string word, return the sum of the number of vowels ('a', 'e', 'i',
// 'o', and 'u') in every substring of word.
// A substring is a contiguous (non-empty) sequence of characters within a string.
// Note: Due to the large constraints, the answer may not fit in a signed 32-bit
// integer. Please be careful during the calculations.
//	Example 1:
//		Input: word = "aba"
//		Output: 6
//		Explanation:
//			All possible substrings are: "a", "ab", "aba", "b", "ba", and "a".
//			- "b" has 0 vowels in it
//			- "a", "ab", "ba", and "a" have 1 vowel each
//			- "aba" has 2 vowels in it
//			Hence, the total sum of vowels = 0 + 1 + 1 + 1 + 1 + 2 = 6.
//	Example 2:
//		Input: word = "abc"
//		Output: 3
//		Explanation:
//			All possible substrings are: "a", "ab", "abc", "b", "bc", and "c".
//			- "a", "ab", and "abc" have 1 vowel each
//			- "b", "bc", and "c" have 0 vowels each
//			Hence, the total sum of vowels = 1 + 1 + 1 + 0 + 0 + 0 = 3.
//	Example 3:
//		Input: word = "ltcd"
//		Output: 0
//		Explanation: There are no vowels in any substring of "ltcd".
//	Example 4:
//		Input: word = "noosabasboosa"
//		Output: 237
//		Explanation: There are a total of 237 vowels in all the substrings.
//	Constraints:
//		1 <= word.length <= 105
//		word consists of lowercase English letters.

// 计算每一个vowel对于整体数字的贡献
// 索引为i，长度为l，则次数为：n = (i+1) * (l-i)
func countVowels(word string) int64 {
	l := len(word)
	var result int64
	for i := range word {
		if word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' {
			result = result + int64((i+1)*(l-i))
		}
	}
	return result
}

// 获取所有的substring，超时
func countVowels2(word string) int64 {
	var result int64
	l := len(word)
	for i := 0; i < l; i++ {
		var vowelCount int64
		if word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' {
			vowelCount++
		}
		result += vowelCount
		for j := i + 1; j < l; j++ {
			if word[j] == 'a' || word[j] == 'e' || word[j] == 'i' || word[j] == 'o' || word[j] == 'u' {
				vowelCount++
			}
			result += vowelCount
		}
	}
	return result
}

func main() {
	fmt.Println(countVowels("aba"))
	fmt.Println(countVowels("abc"))
	fmt.Println(countVowels("ltcd"))
	fmt.Println(countVowels("noosabasboosa"))
}
