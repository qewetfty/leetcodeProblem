package array

import (
	"container/heap"
	"fmt"
)

// You are given a 2D matrix of size m x n, consisting of non-negative integers. You are also given an integer k.
// The value of coordinate (a, b) of the matrix is the XOR of all matrix[i][j] where 0 <= i <= a < m and 0 <= j <= b < n (0-indexed).
// Find the kth largest value (1-indexed) of all the coordinates of matrix.
//	Example 1:
//		Input: matrix = [[5,2],[1,6]], k = 1
//		Output: 7
//		Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the largest value.
//	Example 2:
//		Input: matrix = [[5,2],[1,6]], k = 2
//		Output: 5
//		Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest value.
//	Example 3:
//		Input: matrix = [[5,2],[1,6]], k = 3
//		Output: 4
//		Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd largest value.
//	Example 4:
//		Input: matrix = [[5,2],[1,6]], k = 4
//		Output: 0
//		Explanation: The value of coordinate (1,1) is 5 XOR 2 XOR 1 XOR 6 = 0, which is the 4th largest value.
//	Constraints:
//		m == matrix.length
//		n == matrix[i].length
//		1 <= m, n <= 1000
//		0 <= matrix[i][j] <= 106
//		1 <= k <= m * n

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	h := new(heap1738)
	xorMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		xorMatrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			xorMatrix[i][j] = xorMatrix[i][j] ^ matrix[i][j]
			if i > 0 && j > 0 {
				xorMatrix[i][j] = xorMatrix[i][j] ^ xorMatrix[i-1][j-1] ^ xorMatrix[i-1][j] ^ xorMatrix[i][j-1]
			} else if i > 0 {
				xorMatrix[i][j] = xorMatrix[i][j] ^ xorMatrix[i-1][j]
			} else if j > 0 {
				xorMatrix[i][j] = xorMatrix[i][j] ^ xorMatrix[i][j-1]
			}
			heap.Push(h, xorMatrix[i][j])
			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}
	return heap.Pop(h).(int)
}

type heap1738 []int

func (h heap1738) Len() int {
	return len(h)
}

func (h heap1738) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h heap1738) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heap1738) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heap1738) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func testProblem1738() {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1))
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 2))
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 3))
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 4))
}
