package recursive

import "fmt"

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//	For example, given n = 3, a solution set is:
//		[
//		  "((()))",
//		  "(()())",
//		  "(())()",
//		  "()(())",
//		  "()()()"
//		]

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	generate(0, 0, n, "", &res)
	return res
}

func generate(left int, right int, n int, s string, res *[]string) {
	// terminator
	if left == n && right == n {
		*res = append(*res, s)
		return
	}

	// process
	// drill down
	if left < n {
		generate(left+1, right, n, s+"(", res)
	}
	if right < left {
		generate(left, right+1, n, s+")", res)
	}

	// reverse states
}

func testProblem22() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(2))
}
