package array

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Given a list of non negative integers, arrange them such that they form the largest number.
//	Example 1:
//		Input: [10,2]
//		Output: "210"
//	Example 2:
//		Input: [3,30,34,5,9]
//		Output: "9534330"
//	Note: The result may be very large, so you need to return a string instead of an integer.

func largestNumber(nums []int) string {
	numberList := numbers{num: nums}
	sort.Sort(numberList)
	if numberList.num[len(numberList.num)-1] == 0 {
		return "0"
	}
	resStr := ""
	for i := numberList.Len() - 1; i >= 0; i-- {
		resStr = resStr + strconv.Itoa(numberList.num[i])
	}
	return resStr
}

type numbers struct {
	num []int
}

func (n numbers) Len() int {
	return len(n.num)
}

func (n numbers) Less(i, j int) bool {
	num1, num2 := n.num[i], n.num[j]
	str1, str2 := strconv.Itoa(num1), strconv.Itoa(num2)
	return strings.Compare(str1+str2, str2+str1) <= 0
}

func (n numbers) Swap(i, j int) {
	n.num[i], n.num[j] = n.num[j], n.num[i]
}

func testProblem179() {
	fmt.Println(largestNumber([]int{824, 938, 1399, 5607, 6973, 5703, 8247, 9609, 4398}))
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{0, 0}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
