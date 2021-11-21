package math

import (
	"fmt"
	"strconv"
)

// A k-mirror number is a positive integer without leading zeros that reads the
// same both forward and backward in base-10 as well as in base-k. For example, 9
// is a 2-mirror number. The representation of 9 in base-10 and base-2 are 9 and
// 1001 respectively, which read the same both forward and backward. On the
// contrary, 4 is not a 2-mirror number. The representation of 4 in base-2 is
// 100, which does not read the same both forward and backward.
// Given the base k and the number n, return the sum of the n smallest k-mirror numbers.
//	Example 1:
//		Input: k = 2, n = 5
//		Output: 25
//		Explanation:
//			The 5 smallest 2-mirror numbers and their representations in base-2 are listed as follows:
//			  base-10    base-2
//			    1          1
//			    3          11
//			    5          101
//			    7          111
//			    9          1001
//			Their sum = 1 + 3 + 5 + 7 + 9 = 25.
//	Example 2:
//		Input: k = 3, n = 7
//		Output: 499
//		Explanation:
//			The 7 smallest 3-mirror numbers are and their representations in base-3 are listed as follows:
//			  base-10    base-3
//			    1          1
//			    2          2
//			    4          11
//			    8          22
//			    121        11111
//			    151        12121
//			    212        21212
//			Their sum = 1 + 2 + 4 + 8 + 121 + 151 + 212 = 499.
//	Example 3:
//		Input: k = 7, n = 17
//		Output: 20379000
//		Explanation: The 17 smallest 7-mirror numbers are:
//		1, 2, 3, 4, 5, 6, 8, 121, 171, 242, 292, 16561, 65656, 2137312, 4602064, 6597956, 6958596
//	Constraints:
//		2 <= k <= 9
//		1 <= n <= 30

func kMirror(k int, n int) int64 {
	var (
		result int64
		num    int64
		count  int
	)
	bit := int64(k)
	for count != n {
		num = next10Mirror(num)
		kString := convertToKString(num, bit)
		if isMirror(kString) {
			result += num
			count++
		}
	}
	return result
}

// 枚举下一个回文10进制，分长度为偶数和奇数讨论
// 我们从中间向两边扩散进行寻找。例如 111 -> 121，1111 -> 1221
// 如果遇到9，则向两边扩散进行变换。例如 191 -> 202, 1991 -> 2002
// 如果全是9，则直接加2。例如：9->11，99 -> 101，999 -> 1001
func next10Mirror(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	sByte := []byte(s)
	l := len(s)
	for i := l / 2; i >= 0; i-- {
		if sByte[i] != '9' {
			sByte[i]++
			if l-i-1 != i {
				sByte[l-i-1]++
			}
			for j := i + 1; j <= l/2; j++ {
				sByte[j] = '0'
				sByte[l-j-1] = '0'
			}
			result, _ := strconv.ParseInt(string(sByte), 10, 64)
			return result
		}
	}
	// 全是9，例如9999
	return n + 2
}

func isMirror(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func convertToKString(n int64, k int64) string {
	resByte := make([]byte, 0)
	for n != 0 {
		num := n % k
		resByte = append(resByte, []byte(strconv.FormatInt(num, 10))...)
		n = n / k
	}
	return string(resByte)
}

func testProblem2081() {
	fmt.Println(kMirror(2, 5))
	fmt.Println(kMirror(3, 7))
	fmt.Println(kMirror(7, 17))
}
