package hashTable

import "fmt"

// All DNA is composed of a series of nucleotides abbreviated as 'A', 'C', 'G',
// and 'T', for example: "ACGAATTCCG". When studying DNA, it is sometimes useful
// to identify repeated sequences within the DNA. Write a function to find all
// the 10-letter-long sequences (substrings) that occur more than once in a DNA
// molecule.
//	Example 1:
//		Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//		Output: ["AAAAACCCCC","CCCCCAAAAA"]
//	Example 2:
//		Input: s = "AAAAAAAAAAAAA"
//		Output: ["AAAAAAAAAA"]
//	Constraints:
//		0 <= s.length <= 105
//		s[i] is 'A', 'C', 'G', or 'T'.

func findRepeatedDnaSequences(s string) []string {
	l, res := len(s), make([]string, 0)
	if l < 10 {
		return res
	}
	stringMap := make(map[string]int)
	for i := 0; i < l-9; i++ {
		subStr := s[i : i+10]
		if stringMap[subStr] == 1 {
			res = append(res, subStr)
			stringMap[subStr]++
		} else {
			stringMap[subStr]++
		}
	}
	return res
}

func testProblem187() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
