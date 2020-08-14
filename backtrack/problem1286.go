package backtrack

import "fmt"

// Design an Iterator class, which has:
//	A constructor that takes a string characters of sorted distinct lowercase English letters and a number
// 		combinationLength as arguments.
//	A function next() that returns the next combination of length combinationLength in lexicographical order.
//	A function hasNext() that returns True if and only if there exists a next combination.
//	Example:
//		CombinationIterator iterator = new CombinationIterator("abc", 2); // creates the iterator.
//		iterator.next(); // returns "ab"
//		iterator.hasNext(); // returns true
//		iterator.next(); // returns "ac"
//		iterator.hasNext(); // returns true
//		iterator.next(); // returns "bc"
//		iterator.hasNext(); // returns false
//	Constraints:
//		1 <= combinationLength <= characters.length <= 15
//		There will be at most 10^4 function calls per test.
//		It's guaranteed that all calls of the function next are valid.

type CombinationIterator struct {
	combines []string
	index    int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combines := make([]string, 0)
	backtrack1286(characters, 0, combinationLength, "", &combines)
	return CombinationIterator{
		combines: combines,
		index:    0,
	}
}

func backtrack1286(characters string, start, combinationLength int, currStr string, res *[]string) {
	if len(currStr) == combinationLength {
		*res = append(*res, currStr)
		return
	}
	for i := start; i < len(characters); i++ {
		currStr = currStr + characters[i:i+1]
		backtrack1286(characters, i+1, combinationLength, currStr, res)
		currStr = currStr[:len(currStr)-1]
	}
}

func (this *CombinationIterator) Next() string {
	res := this.combines[this.index]
	this.index++
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return this.index != len(this.combines)
}

func testProblem1286() {
	iterator := Constructor("abc", 2)
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
}
