package greedy

import "fmt"

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//	Symbol       Value
//	I             1
//	V             5
//	X             10
//	L             50
//	C             100
//	D             500
//	M             1000
// For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which
// is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
// Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
// The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//	I can be placed before V (5) and X (10) to make 4 and 9.
//	X can be placed before L (50) and C (100) to make 40 and 90.
//	C can be placed before D (500) and M (1000) to make 400 and 900.
// Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
//	Example 1:
//		Input: 3
//		Output: "III"
//	Example 2:
//		Input: 4
//		Output: "IV"
//	Example 3:
//		Input: 9
//		Output: "IX"
//	Example 4:
//		Input: 58
//		Output: "LVIII"
//		Explanation: L = 50, V = 5, III = 3.
//	Example 5:
//		Input: 1994
//		Output: "MCMXCIV"
//		Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

var romanMap = []map[int]string{
	{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"},
	{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC"},
	{1: "C", 2: "CC", 3: "CCC", 4: "CD", 5: "D", 6: "DC", 7: "DCC", 8: "DCCC", 9: "CM"},
	{1: "M", 2: "MM", 3: "MMM"},
}

// map method
func intToRomanMapFunc(num int) string {
	res, index := "", 0
	for num != 0 {
		curNum := num % 10
		res = romanMap[index][curNum] + res
		num, index = num/10, index+1
	}
	return res
}

var (
	numbers    = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	characters = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// list method
func intToRoman(num int) string {
	res := ""
	for i := 0; i < len(numbers) && num > 0; i++ {
		for numbers[i] <= num && num > 0 {
			res = res + characters[i]
			num = num - numbers[i]
		}
	}
	return res
}

func testProblem12() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
