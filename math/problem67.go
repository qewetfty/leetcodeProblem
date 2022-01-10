package math

import "math/big"

//Given two binary strings a and b, return their sum as a binary string.
//
//
// Example 1:
// Input: a = "11", b = "1"
//Output: "100"
// Example 2:
// Input: a = "1010", b = "1011"
//Output: "10101"
//
//
// Constraints:
//
//
// 1 <= a.length, b.length <= 10⁴
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the zero itself.
//
// Related Topics Math String Bit Manipulation Simulation 👍 3931 👎 467

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	na, _ := big.NewInt(0).SetString(a, 2)
	nb, _ := big.NewInt(0).SetString(b, 2)
	na = na.Add(na, nb)
	return na.Text(2)
}

//leetcode submit region end(Prohibit modification and deletion)
