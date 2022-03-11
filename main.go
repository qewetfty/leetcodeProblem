package main

import (
	"fmt"
	"strings"
)

func repeatLimitedString(s string, repeatLimit int) string {
	l := len(s)
	charList := make([]int, 26)
	for i := 0; i < l; i++ {
		charList[s[i]-'a']++
	}
	result := ""
	for i := 'z'; i >= 'a'; i-- {
		j := i - 'a'
		// 先添加最大的某个字符
		for charList[j] > 0 {
			if charList[j] > repeatLimit {
				charNum := Min(charList[j], repeatLimit)
				charList[j] = charList[j] - charNum
				result = result + strings.Repeat(string(i), charNum)
				// 向某个字符中间插入比它小的字符。如果没有，则直接返回结果
				hasResult := false
				for k := i - 1; k >= 'a'; k-- {
					if charList[k-'a'] == 0 {
						continue
					}
					result = result + string(k)
					charList[k-'a'] = charList[k-'a'] - 1
					hasResult = true
					break
				}
				if !hasResult {
					return result
				}
			} else {
				result = result + strings.Repeat(string(i), charList[j])
				charList[j] = 0
			}
		}
	}
	return result
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	fmt.Println(repeatLimitedString("aababab", 2))
	fmt.Println(repeatLimitedString("cczazcc", 3))
}
