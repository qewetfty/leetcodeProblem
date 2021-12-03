package greedy

//You are given a 0-indexed string street. Each character in street is either
//'H' representing a house or '.' representing an empty space.
//
// You can place buckets on the empty spaces to collect rainwater that falls
//from the adjacent houses. The rainwater from a house at index i is collected if a
//bucket is placed at index i - 1 and/or index i + 1. A single bucket, if placed
//adjacent to two houses, can collect the rainwater from both houses.
//
// Return the minimum number of buckets needed so that for every house, there
//is at least one bucket collecting rainwater from it, or -1 if it is impossible.
//
//
// Example 1:
//
//
//Input: street = "H..H"
//Output: 2
//Explanation:
//We can put buckets at index 1 and index 2.
//"H..H" -> "HBBH" ('B' denotes where a bucket is placed).
//The house at index 0 has a bucket to its right, and the house at index 3 has
//a bucket to its left.
//Thus, for every house, there is at least one bucket collecting rainwater from
//it.
//
//
// Example 2:
//
//
//Input: street = ".H.H."
//Output: 1
//Explanation:
//We can put a bucket at index 2.
//".H.H." -> ".HBH." ('B' denotes where a bucket is placed).
//The house at index 1 has a bucket to its right, and the house at index 3 has
//a bucket to its left.
//Thus, for every house, there is at least one bucket collecting rainwater from
//it.
//
//
// Example 3:
//
//
//Input: street = ".HHH."
//Output: -1
//Explanation:
//There is no empty space to place a bucket to collect the rainwater from the
//house at index 2.
//Thus, it is impossible to collect the rainwater from all the houses.
//
//
// Example 4:
//
//
//Input: street = "H"
//Output: -1
//Explanation:
//There is no empty space to place a bucket.
//Thus, it is impossible to collect the rainwater from the house.
//
//
// Example 5:
//
//
//Input: street = "."
//Output: 0
//Explanation:
//There is no house to collect water from.
//Thus, 0 buckets are needed.
//
//
//
// Constraints:
//
//
// 1 <= street.length <= 10⁵
// street[i] is either'H' or '.'.
//
// Related Topics String Dynamic Programming Greedy 👍 161 👎 7

//leetcode submit region begin(Prohibit modification and deletion)
func minimumBuckets(street string) int {
	l := len(street)
	streetByte := []byte(street)
	result := 0
	for i, char := range streetByte {
		// 如果是"."，遍历下一个
		if char == '.' || char == 'B' {
			continue
		}
		// 如果是H，判断前后是否都是H，若都是，则无法完成
		left, right := byte('H'), byte('H')
		if i > 0 {
			left = streetByte[i-1]
		}
		if i < l-1 {
			right = streetByte[i+1]
		}
		if left == 'H' && right == 'H' {
			return -1
		} else if left == 'H' && right == '.' {
			streetByte[i+1] = 'B'
			result++
			continue
		}
		// 若左边已经被前面覆盖，变为B，则继续执行
		if left == 'B' {
			continue
		}
		if left == '.' && right == 'H' {
			streetByte[i-1] = 'B'
			result++
			continue
		} else if left == '.' && right == '.' {
			streetByte[i+1] = 'B'
			result++
			continue
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
