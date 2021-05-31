package binarySearch

import (
	"fmt"
	"sort"
	"strings"
)

// Given an array of strings products and a string searchWord. We want to design
// a system that suggests at most three product names from products after each
// character of searchWord is typed. Suggested products should have common prefix
// with the searchWord. If there are more than three products with a common
// prefix return the three lexicographically minimums products.
// Return list of lists of the suggested products after each character of searchWord is typed.
//	Example 1:
//		Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
//		Output: [
//			["mobile","moneypot","monitor"],
//			["mobile","moneypot","monitor"],
//			["mouse","mousepad"],
//			["mouse","mousepad"],
//			["mouse","mousepad"]
//			]
//		Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"]
//					 After typing m and mo all products match and we show user ["mobile","moneypot","monitor"]
//					 After typing mou, mous and mouse the system suggests ["mouse","mousepad"]
//	Example 2:
//		Input: products = ["havana"], searchWord = "havana"
//		Output: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
//	Example 3:
//		Input: products = ["bags","baggage","banner","box","cloths"], searchWord = "bags"
//		Output: [["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]
//	Example 4:
//		Input: products = ["havana"], searchWord = "tatiana"
//		Output: [[],[],[],[],[],[],[]]
//	Constraints:
//		1 <= products.length <= 1000
//		There are no repeated elements in products.
//		1 <= Σ products[i].length <= 2 * 10^4
//		All characters of products[i] are lower-case English letters.
//		1 <= searchWord.length <= 1000
//		All characters of searchWord are lower-case English letters.

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	result := make([][]string, 0)
	l := len(searchWord)
	for i := 1; i <= l; i++ {
		result = append(result, searchWords(products, searchWord[:i]))
	}
	return result
}

func searchWords(words []string, word string) []string {
	l := len(words)
	left := binarySearchString(words, word)
	result := make([]string, 0)
	for left < l && len(result) < 3 && strings.HasPrefix(words[left], word) {
		result = append(result, words[left])
		left++
	}
	return result
}

func binarySearchString(words []string, word string) int {
	lo, hi := 0, len(words)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		cmp := strings.Compare(words[mid], word)
		if cmp < 0 {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func testProblem1268() {
	fmt.Println(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
	fmt.Println(suggestedProducts([]string{"havana"}, "havana"))
	fmt.Println(suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"))
	fmt.Println(suggestedProducts([]string{"havana"}, "tatiana"))
}
