package greedy

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are given a string s consisting only of the characters '0' and '1'. In one
// operation, you can change any '0' to '1' or vice versa.
// The string is called alternating if no two adjacent characters are equal. For
// example, the string "010" is alternating, while the string "0100" is not.
// Return the minimum number of operations needed to make s alternating.
//	Example 1:
//		Input: s = "0100"
//		Output: 1
//		Explanation: If you change the last character to '1', s will be "0101", which is alternating.
//	Example 2:
//		Input: s = "10"
//		Output: 0
//		Explanation: s is already alternating.
//	Example 3:
//		Input: s = "1111"
//		Output: 2
//		Explanation: You need two operations to reach "0101" or "1010".
//	Constraints:
//		1 <= s.length <= 104
//		s[i] is either '0' or '1'.

func minOperations1758(s string) int {
	zero, one, l := 0, 0, len(s)
	var zeroLastStr uint8
	zeroLastStr = '0'
	// 先计算s[0]是0的
	if s[0] == '1' {
		zero++
	} else {
		one++
	}
	for i := 1; i < l; i++ {
		if s[i] == zeroLastStr {
			zero++
		} else {
			one++
		}
		if zeroLastStr == '0' {
			zeroLastStr = '1'
		} else {
			zeroLastStr = '0'
		}
	}
	return utils.Min(zero, one)
}

func testProblem1758() {
	fmt.Println(minOperations1758("0100"))
	fmt.Println(minOperations1758("10"))
	fmt.Println(minOperations1758("1111"))
}
