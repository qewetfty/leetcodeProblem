package hashTable

import "fmt"

// Given two strings s and t which consist of only lowercase letters.
// String t is generated by random shuffling string s and then add one more letter at a random position.
// Find the letter that was added in t.
//	Example:
//		Input:
//			s = "abcd"
//			t = "abcde"
//		Output:
//			e
//		Explanation:
//			'e' is the letter that was added.

// hash table method
func findTheDifference(s string, t string) byte {
	charMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		char := t[i]
		if charMap[char] == 0 {
			return char
		}
		charMap[char]--
	}
	return '0'
}

// byte add method
func findTheDifference2(s string, t string) byte {
	sum := 0
	for _, char := range t {
		sum = sum + int(char)
	}
	for _, char := range s {
		sum = sum - int(char)
	}
	return byte(sum)
}

func testProblem389() {
	fmt.Println(findTheDifference2("abcd", "abcde"))
}
