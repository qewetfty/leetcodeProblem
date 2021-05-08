package math

import "strconv"

// Let's say a positive integer is a super-palindrome if it is a palindrome, and
// it is also the square of a palindrome.
// Given two positive integers left and right represented as strings, return the
// number of super-palindromes integers in the inclusive range [left, right].
//	Example 1:
//		Input: left = "4", right = "1000"
//		Output: 4
//		Explanation: 4, 9, 121, and 484 are superpalindromes.
//			Note that 676 is not a superpalindrome: 26 * 26 = 676, but 26 is not a palindrome.
//	Example 2:
//		Input: left = "1", right = "2"
//		Output: 1
//	Constraints:
//		1 <= left.length, right.length <= 18
//		left and right consist of only digits.
//		left and right cannot have leading zeros.
//		left and right represent integers in the range [1, 1018].
//		left is less than or equal to right.

// 回文枚举法
func superpalindromesInRange2(left string, right string) int {
	l, _ := strconv.ParseInt(left, 10, 64)
	r, _ := strconv.ParseInt(right, 10, 64)
	result := 0
	m := 100000
	// 计算奇数
	for i := 1; i < m; i++ {
		s := []byte(strconv.Itoa(i))
		for j := len(s) - 2; j >= 0; j-- {
			s = append(s, s[j])
		}
		n, _ := strconv.ParseInt(string(s), 10, 64)
		n *= n
		if n > r {
			break
		}
		if n >= l && isPalindrome(n) {
			result++
		}
	}
	// 计算偶数
	for i := 1; i < m; i++ {
		s := []byte(strconv.Itoa(i))
		for j := len(s) - 1; j >= 0; j-- {
			s = append(s, s[j])
		}
		n, _ := strconv.ParseInt(string(s), 10, 64)
		n *= n
		if n > r {
			break
		}
		if n >= l && isPalindrome(n) {
			result++
		}
	}
	return result
}

func isPalindrome(n int64) bool {
	return n == reverseLong(n)
}

func reverseLong(n int64) int64 {
	var result int64
	result = 0
	for n > 0 {
		result = 10*result + n%10
		n = n / 10
	}
	return result
}

// 暴力打表法
var resultSet = []int64{1, 4, 9, 121, 484, 10201, 12321, 14641, 40804, 44944, 1002001, 1234321, 4008004, 100020001,
	102030201, 104060401, 121242121, 123454321, 125686521, 400080004, 404090404, 10000200001, 10221412201, 12102420121,
	12345654321, 40000800004, 1000002000001, 1002003002001, 1004006004001, 1020304030201, 1022325232201, 1024348434201,
	1210024200121, 1212225222121, 1214428244121, 1232346432321, 1234567654321, 4000008000004, 4004009004004,
	100000020000001, 100220141022001, 102012040210201, 102234363432201, 121000242000121, 121242363242121, 123212464212321,
	123456787654321, 400000080000004, 10000000200000001, 10002000300020001, 10004000600040001, 10020210401202001,
	10022212521222001, 10024214841242001, 10201020402010201, 10203040504030201, 10205060806050201, 10221432623412201,
	10223454745432201, 12100002420000121, 12102202520220121, 12104402820440121, 12122232623222121, 12124434743442121,
	12321024642012321, 12323244744232321, 12343456865434321, 12345678987654321, 40000000800000004, 40004000900040004,
	1000000002000000001, 1000220014100220001, 1002003004003002001, 1002223236323222001, 1020100204020010201, 1020322416142230201,
	1022123226223212201, 1022345658565432201, 1210000024200000121, 1210242036302420121, 1212203226223022121, 1212445458545442121,
	1232100246420012321, 1232344458544432321, 1234323468643234321, 4000000008000000004}

func superpalindromesInRange(left string, right string) int {
	l, _ := strconv.ParseInt(left, 10, 64)
	r, _ := strconv.ParseInt(right, 10, 64)
	result := 0
	for _, num := range resultSet {
		if num >= l && num <= r {
			result++
		}
	}
	return result
}
