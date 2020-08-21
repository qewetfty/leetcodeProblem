package array

import "fmt"

// Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
// You may return any answer array that satisfies this condition.
//	Example 1:
//		Input: [3,1,2,4]
//		Output: [2,4,3,1]
//		The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
//	Note:
//		1 <= A.length <= 5000
//		0 <= A[i] <= 5000

// extra array method
func sortArrayByParity2(A []int) []int {
	odds, evens := make([]int, 0), make([]int, 0)
	for _, num := range A {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	evens = append(evens, odds...)
	return evens
}

// no extra array method
func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		if A[i]%2 == 0 {
			i++
		} else {
			A[i], A[j] = A[j], A[i]
			j--
		}
	}
	return A
}

func testProblem905() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
