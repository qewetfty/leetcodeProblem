package string

import "strings"

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return sLen
	}
	sByte := []byte(s)
	sSub := make([]byte, 0)
	longest := 0
	for i := 0; i < sLen; i++ {
		nextString := sByte[i]
		nextIndex := strings.Index(string(sSub), string(nextString))
		if nextIndex == -1 {
			sSub = append(sSub, nextString)
		} else {
			sSub = sSub[nextIndex+1:]
			sSub = append(sSub, nextString)
		}
		if longest < len(sSub) {
			longest = len(sSub)
		}
	}
	return longest
}
