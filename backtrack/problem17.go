package backtrack

import "fmt"

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//	Example:
//		Input: "23"
//		Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//	Note:
//		Although the above answer is in lexicographical order, your answer could be in any order you want.

var letterMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	backtrack17(digits, 0, "", &res)
	return res
}

func backtrack17(digit string, start int, letter string, res *[]string) {
	if len(letter) == len(digit) {
		*res = append(*res, letter)
		return
	}
	for i := start; i < len(digit); i++ {
		num := digit[i : i+1]
		for _, l := range letterMap[num] {
			letter = letter + l
			backtrack17(digit, i+1, letter, res)
			letter = letter[:len(letter)-1]
		}
	}
}

func testProblem17() {
	fmt.Println(letterCombinations("23"))
}
