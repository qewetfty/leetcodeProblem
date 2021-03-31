package greedy

import (
	"bytes"
	"fmt"
)

// You want to form a target string of lowercase letters.
// At the beginning, your sequence is target.length '?' marks.  You also have a stamp of lowercase letters.
// On each turn, you may place the stamp over the sequence, and replace every
// letter in the sequence with the corresponding letter from the stamp. You can
// make up to 10 * target.length turns.
// For example, if the initial sequence is "?????", and your stamp is "abc", then
// you may make "abc??", "?abc?", "??abc" in the first turn. (Note that the stamp
// must be fully contained in the boundaries of the sequence in order to stamp.)
// If the sequence is possible to stamp, then return an array of the index of the
// left-most letter being stamped at each turn. If the sequence is not possible
// to stamp, return an empty array.
// For example, if the sequence is "ababc", and the stamp is "abc", then we could
// return the answer [0, 2], corresponding to the moves "?????" -> "abc??" ->
// "ababc".
// Also, if the sequence is possible to stamp, it is guaranteed it is possible to
// stamp within 10 * target.length moves. Any answers specifying more than this
// number of moves will not be accepted.
//	Example 1:
//		Input: stamp = "abc", target = "ababc"
//		Output: [0,2]
//		([1,0,2] would also be accepted as an answer, as well as some other answers.)
//	Example 2:
//		Input: stamp = "abca", target = "aabcaca"
//		Output: [3,0,1]
//	Note:
//		1 <= stamp.length <= target.length <= 1000
//		stamp and target only contain lowercase letters.

func movesToStamp(stamp string, target string) []int {
	sLen, tLen := len(stamp), len(target)
	bTarget, bStamp := make([]byte, tLen), []byte(stamp)
	for i := 0; i < tLen; i++ {
		bTarget[i] = '?'
	}
	now := []byte(target)
	ret := make([]int, 0)
	flag := true
	for flag {
		flag = false
		// 如果都是'?'了说明找完了，写结果
		if bytes.Compare(bTarget, now) == 0 {
			retLen := len(ret)
			for i := 0; i < retLen/2; i++ {
				ret[i], ret[retLen-i-1] = ret[retLen-i-1], ret[i]
			}
			return ret
		}
		// 从头开始匹配
		for i := 0; i <= tLen-sLen; i++ {
			// 如果全是'?'，不用匹配
			if bytes.Compare(bTarget[i:i+sLen], now[i:i+sLen]) == 0 {
				continue
			}
			if compare(bStamp, now[i:i+sLen]) {
				ret = append(ret, i)
				for j := i; j < i+sLen; j++ {
					now[j] = '?'
				}
				flag = true
				break
			}
		}
	}
	return []int{}
}

func compare(src, tar []byte) bool {
	if bytes.Compare(src, tar) == 0 {
		return true
	}
	l := len(src)
	for i := 0; i < l; i++ {
		if src[i] != tar[i] && tar[i] != '?' {
			return false
		}
	}
	return true
}

func testProblem936() {
	fmt.Println(movesToStamp("abc", "ababc"))
	fmt.Println(movesToStamp("abca", "aabcaca"))
}
