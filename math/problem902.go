package math

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// Given an array of digits, you can write numbers using each digits[i] as many
// times as we want. For example, if digits = ['1','3','5'], we may write numbers
// such as '13', '551', and '1351315'.
// Return the number of positive integers that can be generated that are less than or equal to a given integer n.
//	Example 1:
//		Input: digits = ["1","3","5","7"], n = 100
//		Output: 20
//		Explanation:
//			The 20 numbers that can be written are:
//			1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
//	Example 2:
//		Input: digits = ["1","4","9"], n = 1000000000
//		Output: 29523
//		Explanation:
//			We can write 3 one digit numbers, 9 two digit numbers, 27 three digit numbers,
//			81 four digit numbers, 243 five digit numbers, 729 six digit numbers,
//			2187 seven digit numbers, 6561 eight digit numbers, and 19683 nine digit numbers.
//			In total, this is 29523 integers that can be written using the digits array.
//	Example 3:
//		Input: digits = ["7"], n = 8
//		Output: 1
//	Constraints:
//		1 <= digits.length <= 9
//		digits[i].length == 1
//		digits[i] is a digit from '1' to '9'.
//		All the values in digits are unique.
//		1 <= n <= 109

func atMostNGivenDigitSet(digits []string, n int) int {
	nums := make([]int, 0)
	// 把字符串转换为数字
	for _, digit := range digits {
		num, _ := strconv.Atoi(digit)
		nums = append(nums, num)
	}
	sort.Ints(nums)
	// n比第一个数还小，直接返回
	if n < nums[0] {
		return 0
	}
	// 分解n的各个位，存入slice，方便之后按位比对结果
	x, numberN := n, make([]int, 0)
	for x > 0 {
		numberN = append(numberN, x%10)
		x /= 10
	}

	// 首先把位数比n低的数全部算上，累加求和
	res, numberLen, digitLen := 0, len(numberN), len(nums)
	for i := 1; i < numberLen; i++ {
		res = res + int(math.Pow(float64(digitLen), float64(i)))
	}
	// 遍历n的各个位，numberN倒序存储，所以从后向前遍历
	start := numberLen - 1
	for start >= 0 {
		curNum, findEqual := numberN[start], false
		for _, num := range nums {
			if curNum < num {
				return res
				// 如果找到了相等的才往下一位判断，如果没找到相等的，直接return
			} else if curNum == num {
				if start == 0 {
					res++
				}
				findEqual = true
				break
			} else {
				res = res + int(math.Pow(float64(digitLen), float64(start)))
			}
		}
		if !findEqual {
			break
		}
		start--
	}
	return res
}

func testProblem902() {
	fmt.Println(atMostNGivenDigitSet([]string{"1"}, 834))
	fmt.Println(atMostNGivenDigitSet([]string{"9"}, 55))
	fmt.Println(atMostNGivenDigitSet([]string{"1", "2", "4"}, 247))
	fmt.Println(atMostNGivenDigitSet([]string{"7"}, 8))

	fmt.Println(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
	fmt.Println(atMostNGivenDigitSet([]string{"1", "4", "9"}, 1000000000))
	fmt.Println(atMostNGivenDigitSet([]string{"1", "3", "4", "8", "9"}, 75205013))
}
