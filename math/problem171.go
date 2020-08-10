package math

import (
	"fmt"
	"math"
)

// Given a column title as appear in an Excel sheet, return its corresponding column number.
//	For example:
//    	A -> 1
//    	B -> 2
//    	C -> 3
//    	...
//    	Z -> 26
//    	AA -> 27
//    	AB -> 28
//    	...
//	Example 1:
//		Input: "A"
//		Output: 1
//	Example 2:
//		Input: "AB"
//		Output: 28
//	Example 3:
//		Input: "ZY"
//		Output: 701
//	Constraints:
//		1 <= s.length <= 7
//		s consists only of uppercase English letters.
//		s is between "A" and "FXSHRXW".

func titleToNumber(s string) int {
	l := len(s)
	res := 0
	for i := 0; i < l; i++ {
		res = res + int(s[i]-'A'+1)*int(math.Pow(26.0, float64(l-i-1)))
	}
	return res
}

func titleToNumber2(s string) int {
	l := len(s)
	res := 0
	for i := 0; i < l; i++ {
		res = res*26 + int(s[i]-'A'+1)
	}
	return res
}

func testProblem171() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
	fmt.Println(titleToNumber("FXSHRXW"))
	fmt.Println(titleToNumber("AA"))
}
