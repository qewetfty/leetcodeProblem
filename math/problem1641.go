package math

import "fmt"

// Given an integer n, return the number of strings of length n that consist only
// of vowels (a, e, i, o, u) and are lexicographically sorted.
// A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes before s[i+1] in the alphabet.
//	Example 1:
//		Input: n = 1
//		Output: 5
//		Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].
//	Example 2:
//		Input: n = 2
//		Output: 15
//		Explanation: The 15 sorted strings that consist of vowels only are
//		["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
//		Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.
//	Example 3:
//		Input: n = 33
//		Output: 66045
//	Constraints:
//		1 <= n <= 50

// 排列组合解法，相当于把n个小球放入m个盒子中，盒子可以为空的方法数。这里m=5
// 参考解析：https://leetcode-cn.com/problems/count-sorted-vowel-strings/solution/zhong-xue-shu-xue-ke-pu-n-ge-xiao-qiu-fang-dao-m-g/
func countVowelStrings(n int) int {
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}

// 回溯写法
func countVowelStrings2(n int) int {
	return backtrack(n, 1, 0)
}

func backtrack(n, i int, char int) int {
	if i == n {
		return 5 - char
	}
	number := 0
	for j := char; j < 5; j++ {
		number = number + backtrack(n, i+1, j)
	}
	return number
}

func testProblem1641() {
	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
	fmt.Println(countVowelStrings(3))
	fmt.Println(countVowelStrings(4))
	fmt.Println(countVowelStrings(5))
	fmt.Println(countVowelStrings(6))
	fmt.Println(countVowelStrings(7))
}
