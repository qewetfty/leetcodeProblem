package math

//Given a positive integer num, return the number of positive integers less
//than or equal to num whose digit sums are even.
//
// The digit sum of a positive integer is the sum of all its digits.
//
//
// Example 1:
//
//
//Input: num = 4
//Output: 2
//Explanation:
//The only integers less than or equal to 4 whose digit sums are even are 2 and
//4.
//
//
// Example 2:
//
//
//Input: num = 30
//Output: 14
//Explanation:
//The 14 integers less than or equal to 30 whose digit sums are even are
//2, 4, 6, 8, 11, 13, 15, 17, 19, 20, 22, 24, 26, and 28.
//
//
//
// Constraints:
//
//
// 1 <= num <= 1000
//
// Related Topics Math Simulation 👍 126 👎 8

//leetcode submit region begin(Prohibit modification and deletion)
func countEven(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		curNum, curSum := i, 0
		for curNum > 0 {
			curSum = curSum + curNum%10
			curNum = curNum / 10
		}
		if curSum%2 == 0 {
			result++
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
