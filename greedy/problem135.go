package greedy

import "leetcodeProblem/utils"

//There are n children standing in a line. Each child is assigned a rating
//value given in the integer array ratings.
//
// You are giving candies to these children subjected to the following
//requirements:
//
//
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
//
//
// Return the minimum number of candies you need to have to distribute the
//candies to the children.
//
//
// Example 1:
//
//
//Input: ratings = [1,0,2]
//Output: 5
//Explanation: You can allocate to the first, second and third child with 2, 1,
//2 candies respectively.
//
//
// Example 2:
//
//
//Input: ratings = [1,2,2]
//Output: 4
//Explanation: You can allocate to the first, second and third child with 1, 2,
//1 candies respectively.
//The third child gets 1 candy because it satisfies the above two conditions.
//
//
//
// Constraints:
//
//
// n == ratings.length
// 1 <= n <= 2 * 10⁴
// 0 <= ratings[i] <= 2 * 10⁴
//
// Related Topics Array Greedy 👍 2985 👎 261

//leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
	l := len(ratings)
	left := make([]int, l)
	for i := 0; i < l; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	result := 0
	for i := l - 1; i >= 0; i-- {
		if i < l-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		result = result + utils.Max(left[i], right)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
