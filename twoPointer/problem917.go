package twoPointer

// Given a string s, reverse the string according to the following rules:
//	All the characters that are not English letters remain in the same position.
//	All the English letters (lowercase or uppercase) should be reversed.
// Return s after reversing it.
//	Example 1:
//		Input: s = "ab-cd"
//		Output: "dc-ba"
//	Example 2:
//		Input: s = "a-bC-dEf-ghIj"
//		Output: "j-Ih-gfE-dCba"
//	Example 3:
//		Input: s = "Test1ng-Leet=code-Q!"
//		Output: "Qedo1ct-eeLg=ntse-T!"
//	Constraints:
//		1 <= s.length <= 100
//		s consists of characters with ASCII values in the range [33, 122].
//		s does not contain '\"' or '\\'.

func reverseOnlyLetters(s string) string {
	b, l := []byte(s), len(s)
	left, right := 0, l-1
	for left < right {
		for left < l && (b[left] < 'A' || (b[left] > 'Z' && b[left] < 'a') || b[left] > 'z') {
			left++
		}
		for right >= 0 && (b[right] < 'A' || (b[right] > 'Z' && b[right] < 'a') || b[right] > 'z') {
			right--
		}
		if left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}
	return string(b)
}
