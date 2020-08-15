package hashTable

import "fmt"

// Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that
// can be built with those letters.
// This is case sensitive, for example "Aa" is not considered a palindrome here.
//	Note:
//		Assume the length of given string will not exceed 1,010.
//	Example:
//	Input:
//		"abccccdd"
//	Output:
//		7
//	Explanation:
//		One longest palindrome that can be built is "dccaccd", whose length is 7.

func longestPalindrome(s string) int {
	l, res, odd := len(s), 0, false
	if l == 0 {
		return res
	}
	charMap := make(map[byte]int)
	for i := 0; i < l; i++ {
		if _, ok := charMap[s[i]]; ok {
			charMap[s[i]]++
		} else {
			charMap[s[i]] = 1
		}
	}
	for _, appearTime := range charMap {
		if appearTime%2 == 0 {
			res = res + appearTime
		} else {
			res = res + appearTime - 1
			odd = true
		}
	}
	if odd {
		res = res + 1
	}
	return res
}

func testProblem409() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("aabcccddd"))
	fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
