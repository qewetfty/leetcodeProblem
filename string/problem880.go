package string

import (
	"fmt"
	"strings"
)

// An encoded string S is given. To find and write the decoded string to a tape,
// the encoded string is read one character at a time and the following steps are
// taken:
//		If the character read is a letter, that letter is written onto the tape.
//		If the character read is a digit (say d), the entire current tape is repeatedly written d-1 more times in total.
//	Now for some encoded string S, and an index K, find and return the K-th letter (1 indexed) in the decoded string.
//	Example 1:
//		Input: S = "leet2code3", K = 10
//		Output: "o"
//		Explanation:
//			The decoded string is "leetleetcodeleetleetcodeleetleetcode".
//			The 10th letter in the string is "o".
//	Example 2:
//		Input: S = "ha22", K = 5
//		Output: "h"
//		Explanation:
//			The decoded string is "hahahaha".  The 5th letter is "h".
//	Example 3:
//		Input: S = "a2345678999999999999999", K = 1
//		Output: "a"
//		Explanation:
//			The decoded string is "a" repeated 8301530446056247680 times.  The 1st letter is "a".
//	Constraints:
//		2 <= S.length <= 100
//		S will only contain lowercase letters and digits 2 through 9.
//		S starts with a letter.
//		1 <= K <= 10^9
//		It's guaranteed that K is less than or equal to the length of the decoded string.
//		The decoded string is guaranteed to have less than 2^63 letters.

// 只计算最终长度，K%size 和K的结果是一样的，以此去判断最终结果
func decodeAtIndex(S string, K int) string {
	l := len(S)
	size := 0
	for i := 0; i < l; i++ {
		curChar := S[i]
		if isDigit(curChar) {
			size = size * int(S[i]-'0')
		} else {
			size++
		}
	}
	for i := l - 1; i >= 0; i-- {
		curChar := S[i]
		K = K % size
		if K == 0 && !isDigit(curChar) {
			return S[i : i+1]
		}
		if isDigit(curChar) {
			size = size / int(curChar-'0')
		} else {
			size--
		}
	}
	return ""
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

// 使用字符串模拟，超时
func decodeAtIndex2(S string, K int) string {
	l, str := len(S), ""
	for i := 0; i < l; i++ {
		curChar, curStr := S[i], S[i:i+1]
		if curChar >= '0' && curChar <= '9' {
			repeatTime := int(curChar - '0')
			str = strings.Repeat(str, repeatTime)
		} else {
			str = str + curStr
		}
		if len(str) >= K {
			return str[K-1 : K]
		}
	}
	return str[K-1 : K]
}

func testProblem880() {
	fmt.Println(decodeAtIndex("leet2code3", 10))
	fmt.Println(decodeAtIndex("a2345678999999999999999", 1))
	fmt.Println(decodeAtIndex("ha22", 5))
	fmt.Println(decodeAtIndex("yyuele72uthzyoeut7oyku2yqmghy5luy9qguc28ukav7an6a2bvizhph35t86qicv4gyeo6av7gerovv5lnw47954bsv2xruaej", 123365626))
}
