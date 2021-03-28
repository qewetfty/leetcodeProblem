package math

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a non-empty string containing an out-of-order English representation of digits 0-9, output the digits in ascending order.
//	Note:
//		Input contains only lowercase English letters.
//		Input is guaranteed to be valid and can be transformed to its original digits. That means invalid inputs such as "abc" or "zerone" are not permitted.
//		Input length is less than 50,000.
//	Example 1:
//		Input: "owoztneoer"
//		Output: "012"
//	Example 2:
//		Input: "fviefuro"
//		Output: "45"

func originalDigits(s string) string {
	chars, l := make([]int, 26), len(s)
	numberCount := make([]int, 10)
	for i := 0; i < l; i++ {
		chars[s[i]-'a']++
	}
	// 开始处理数字部分
	// 处理0
	zeroCount := chars['z'-'a']
	numberCount[0] = zeroCount
	chars['z'-'a'], chars['e'-'a'], chars['r'-'a'], chars['o'-'a'] = chars['z'-'a']-zeroCount, chars['e'-'a']-zeroCount, chars['r'-'a']-zeroCount, chars['o'-'a']-zeroCount
	// 处理2
	twoCount := chars['w'-'a']
	numberCount[2] = twoCount
	chars['t'-'a'], chars['w'-'a'], chars['o'-'a'] = chars['t'-'a']-twoCount, chars['w'-'a']-twoCount, chars['o'-'a']-twoCount
	// 处理6
	sixCount := chars['x'-'a']
	numberCount[6] = sixCount
	chars['s'-'a'], chars['i'-'a'], chars['x'-'a'] = chars['s'-'a']-sixCount, chars['i'-'a']-sixCount, chars['x'-'a']-sixCount
	// 处理8
	eightCount := chars['g'-'a']
	numberCount[8] = eightCount
	chars['e'-'a'], chars['i'-'a'], chars['g'-'a'], chars['h'-'a'], chars['t'-'a'] = chars['e'-'a']-eightCount, chars['i'-'a']-eightCount, chars['g'-'a']-eightCount, chars['h'-'a']-eightCount, chars['t'-'a']-eightCount
	// 处理7
	sevenCount := chars['s'-'a']
	numberCount[7] = sevenCount
	chars['s'-'a'], chars['e'-'a'], chars['v'-'a'], chars['b'-'a'] = chars['s'-'a']-sevenCount, chars['e'-'a']-sevenCount*2, chars['v'-'a']-sevenCount, chars['h'-'a']-sevenCount
	// 处理4
	fourCount := chars['u'-'a']
	numberCount[4] = fourCount
	chars['f'-'a'], chars['o'-'a'], chars['u'-'a'], chars['r'-'a'] = chars['f'-'a']-fourCount, chars['o'-'a']-fourCount, chars['u'-'a']-fourCount, chars['r'-'a']-fourCount
	// 处理5
	fiveCount := chars['f'-'a']
	numberCount[5] = fiveCount
	chars['f'-'a'], chars['i'-'a'], chars['v'-'a'], chars['e'-'a'] = chars['f'-'a']-fiveCount, chars['i'-'a']-fiveCount, chars['v'-'a']-fiveCount, chars['e'-'a']-fiveCount
	// 处理3
	threeCount := chars['r'-'a']
	numberCount[3] = threeCount
	chars['t'-'a'], chars['h'-'a'], chars['r'-'a'], chars['e'-'a'] = chars['t'-'a']-threeCount, chars['h'-'a']-threeCount, chars['r'-'a']-threeCount, chars['e'-'a']-threeCount*2
	// 处理1
	oneCount := chars['o'-'a']
	numberCount[1] = oneCount
	chars['o'-'a'], chars['n'-'a'], chars['e'-'a'] = chars['o'-'a']-oneCount, chars['n'-'a']-oneCount, chars['e'-'a']-oneCount
	// 处理9
	nineCount := chars['e'-'a']
	numberCount[9] = nineCount
	// 写结果
	result := ""
	for i, num := range numberCount {
		if num > 0 {
			result = result + strings.Repeat(strconv.Itoa(i), num)
		}
	}
	return result
}

func testProblem423() {
	fmt.Println(originalDigits("nnei"))
	fmt.Println(originalDigits("owoztneoer"))
	fmt.Println(originalDigits("fviefuro"))
}
