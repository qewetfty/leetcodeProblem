package greedy

import "leetcodeProblem/utils"

//You are given a string s and an integer repeatLimit. Construct a new string
//repeatLimitedString using the characters of s such that no letter appears more
//than repeatLimit times in a row. You do not have to use all characters from s.
//
// Return the lexicographically largest repeatLimitedString possible.
//
// A string a is lexicographically larger than a string b if in the first
//position where a and b differ, string a has a letter that appears later in the
//alphabet than the corresponding letter in b. If the first min(a.length, b.length)
//characters do not differ, then the longer string is the lexicographically larger one.
//
//
//
// Example 1:
//
//
//Input: s = "cczazcc", repeatLimit = 3
//Output: "zzcccac"
//Explanation: We use all of the characters from s to construct the
//repeatLimitedString "zzcccac".
//The letter 'a' appears at most 1 time in a row.
//The letter 'c' appears at most 3 times in a row.
//The letter 'z' appears at most 2 times in a row.
//Hence, no letter appears more than repeatLimit times in a row and the string
//is a valid repeatLimitedString.
//The string is the lexicographically largest repeatLimitedString possible so
//we return "zzcccac".
//Note that the string "zzcccca" is lexicographically larger but the letter 'c'
//appears more than 3 times in a row, so it is not a valid repeatLimitedString.
//
//
// Example 2:
//
//
//Input: s = "aababab", repeatLimit = 2
//Output: "bbabaa"
//Explanation: We use only some of the characters from s to construct the
//repeatLimitedString "bbabaa".
//The letter 'a' appears at most 2 times in a row.
//The letter 'b' appears at most 2 times in a row.
//Hence, no letter appears more than repeatLimit times in a row and the string
//is a valid repeatLimitedString.
//The string is the lexicographically largest repeatLimitedString possible so
//we return "bbabaa".
//Note that the string "bbabaaa" is lexicographically larger but the letter 'a'
//appears more than 2 times in a row, so it is not a valid repeatLimitedString.
//
//
//
// Constraints:
//
//
// 1 <= repeatLimit <= s.length <= 10⁵
// s consists of lowercase English letters.
//
// Related Topics String Greedy Heap (Priority Queue) Counting 👍 308 👎 20

//leetcode submit region begin(Prohibit modification and deletion)
func repeatLimitedString(s string, repeatLimit int) string {
	l := len(s)
	charList := make([]int, 26)
	var cur byte
	cur = 'a'
	for i := 0; i < l; i++ {
		charList[s[i]-'a']++
		if s[i] > cur {
			cur = s[i]
		}
	}
	result := make([]byte, 0)
	for cur >= 'a' {
		next := cur - 1
		for next >= 'a' && charList[next-'a'] == 0 {
			next--
		}
		num := utils.Min(charList[cur-'a'], repeatLimit)
		for i := 0; i < num; i++ {
			result = append(result, cur)
		}
		charList[cur-'a'] -= num
		if charList[cur-'a'] == 0 {
			cur, next = next, next-1
		} else if next < 'a' {
			break
		} else {
			result = append(result, next)
			charList[next-'a']--
		}
	}
	return string(result)
}
