package dynamic_programming

//Alice is texting Bob using her phone. The mapping of digits to letters is
//shown in the figure below.
//
// In order to add a letter, Alice has to press the key of the corresponding
//digit i times, where i is the position of the letter in the key.
//
//
// For example, to add the letter 's', Alice has to press '7' four times.
//Similarly, to add the letter 'k', Alice has to press '5' twice.
// Note that the digits '0' and '1' do not map to any letters, so Alice does
//not use them.
//
//
// However, due to an error in transmission, Bob did not receive Alice's text
//message but received a string of pressed keys instead.
//
//
// For example, when Alice sent the message "bob", Bob received the string "2266
//622".
//
//
// Given a string pressedKeys representing the string received by Bob, return
//the total number of possible text messages Alice could have sent.
//
// Since the answer may be very large, return it modulo 10⁹ + 7.
//
//
// Example 1:
//
//
//Input: pressedKeys = "22233"
//Output: 8
//Explanation:
//The possible text messages Alice could have sent are:
//"aaadd", "abdd", "badd", "cdd", "aaae", "abe", "bae", and "ce".
//Since there are 8 possible messages, we return 8.
//
//
// Example 2:
//
//
//Input: pressedKeys = "222222222222222222222222222222222222"
//Output: 82876089
//Explanation:
//There are 2082876103 possible text messages Alice could have sent.
//Since we need to return the answer modulo 10⁹ + 7, we return 2082876103 % (10⁹
// + 7) = 82876089.
//
//
//
// Constraints:
//
//
// 1 <= pressedKeys.length <= 10⁵
// pressedKeys only consists of digits from '2' - '9'.
//
// 👍 315 👎 6

//leetcode submit region begin(Prohibit modification and deletion)
var d = 1000000007

func countTexts(pressedKeys string) int {
	l := len(pressedKeys)
	char, count := pressedKeys[0], 1
	result := 1
	for i := 1; i < l; i++ {
		curChar := pressedKeys[i]
		if curChar == char {
			count++
		} else {
			var curCount int
			if char == '7' || char == '9' {
				curCount = getCount2(count)
			} else {
				curCount = getCount(count)
			}
			result = (result * curCount) % d
			char, count = curChar, 1
		}
	}
	var curCount int
	if char == '7' || char == '9' {
		curCount = getCount2(count)
	} else {
		curCount = getCount(count)
	}
	result = (result * curCount) % d
	return result
}

func getCount(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	x, y, z := 1, 2, 4
	for i := 4; i <= n; i++ {
		x, y, z = y, z, (x+y+z)%d
	}
	return z
}

func getCount2(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	if n == 4 {
		return 8
	}
	w, x, y, z := 1, 2, 4, 8
	for i := 5; i <= n; i++ {
		w, x, y, z = x, y, z, (w+x+y+z)%d
	}
	return z
}

//leetcode submit region end(Prohibit modification and deletion)
