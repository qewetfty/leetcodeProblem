package array

import "fmt"

//We have an array A of integers, and an array queries of queries.
//For the i-th query val = queries[i][0], index = queries[i][1],
// we add val to A[index].  Then, the answer to the i-th query is the sum of the even values of A.
//(Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.)
//Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.

func testProblem985() {
	A := []int{1, 2, 3, 4}
	querys := [][]int{[]int{1, 0}, []int{-3, 1}, []int{-4, 0}, []int{2, 3}}
	fmt.Println(sumEvenAfterQueries(A, querys))
	A2 := []int{-1, 3, -3, 9, -6, 8, -5}
	query2 := [][]int{{-5, 1}, {10, 2}, {-6, 3}, {3, 2}, {9, 5}, {7, 5}, {8, 0}}
	fmt.Println(sumEvenAfterQueries(A2, query2))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	evenSum := 0
	evenList := make([]int, len(A))
	result := make([]int, 0)
	for i, num := range A {
		if num%2 == 0 {
			evenList[i] = 1
			evenSum += num
		} else {
			evenList[i] = 0
		}
	}
	for _, op := range queries {
		index := op[1]
		add := op[0]
		if evenList[index] == 1 {
			if add%2 == 0 {
				evenSum += add
			} else {
				evenSum -= A[index]
				evenList[index] = 0
			}
		} else {
			if add%2 == 1 || add%2 == -1 {
				evenSum = evenSum + add + A[index]
				evenList[index] = 1
			}
		}
		A[index] += add
		result = append(result, evenSum)
	}
	return result
}
