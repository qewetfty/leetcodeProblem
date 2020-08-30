package unionFind

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Given a non-empty array of unique positive integers A, consider the following graph:
// There are A.length nodes, labelled A[0] to A[A.length - 1];
// There is an edge between A[i] and A[j] if and only if A[i] and A[j] share a common factor greater than 1.
// Return the size of the largest connected component in the graph.
//	Example 1:
//		Input: [4,6,15,35]
//		Output: 4
//	Example 2:
//		Input: [20,50,9,63]
//		Output: 2
//	Example 3:
//		Input: [2,3,6,7,4,12,21,39]
//		Output: 8
//	Note:
//		1 <= A.length <= 20000
//		1 <= A[i] <= 100000

func largestComponentSize(A []int) int {
	maxVal := 0
	for _, num := range A {
		maxVal = utils.Max(maxVal, num)
	}
	union := newUnionFind(maxVal + 1)
	for _, num := range A {
		upper := int(math.Sqrt(float64(num)))
		for i := 2; i <= upper; i++ {
			if num%i == 0 {
				union.Union(num, i)
				union.Union(num, num/i)
			}
		}
	}
	res, unionMap := 0, make(map[int]int)
	for _, num := range A {
		root, rootNum := union.Parent(num), 0
		if _, ok := unionMap[root]; ok {
			rootNum = unionMap[root]
		}
		rootNum++
		unionMap[root] = rootNum
		res = utils.Max(res, rootNum)
	}
	return res
}

// 一一遍历求每一个数的最大公约数，做一个并查集，求并查集最大root（超时）
func largestComponentSizeExceed(A []int) int {
	l := len(A)
	u := newUnionFind(l)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			g := gcd(A[i], A[j])
			if g != 1 {
				u.Union(i, j)
			}
		}
	}
	parentMap, max := make(map[int]int), 0
	for i := range u.parent {
		root, num := u.Parent(i), 0
		if _, ok := parentMap[root]; ok {
			num = parentMap[root]
		}
		num++
		parentMap[root] = num
		max = utils.Max(max, num)
	}
	return max
}

func gcd(m, n int) int {
	if m < n {
		m, n = n, m
	}
	r := m % n
	for r > 0 {
		m, n = n, r
		r = m % n
	}
	return n
}

type unionFind struct {
	parent []int
	count  int
}

func newUnionFind(n int) *unionFind {
	union := new(unionFind)
	union.parent = make([]int, n)
	for i := 0; i < n; i++ {
		union.parent[i] = i
	}
	union.count = n
	return union
}

func (union *unionFind) Parent(p int) int {
	root, i := p, p
	for union.parent[root] != root {
		root = union.parent[root]
	}
	for union.parent[i] != i {
		i, union.parent[i] = union.parent[i], root
	}
	return root
}

func (union *unionFind) Union(i, j int) {
	p1, p2 := union.Parent(i), union.Parent(j)
	if p1 == p2 {
		return
	}
	union.parent[p1] = p2
	union.count--
}

func (union *unionFind) Count() int {
	return union.count
}

func testProblem952() {
	fmt.Println(largestComponentSize([]int{4, 6, 15, 35}))
	fmt.Println(largestComponentSize([]int{20, 50, 9, 63}))
	fmt.Println(largestComponentSize([]int{2, 3, 6, 7, 4, 12, 21, 39}))
	fmt.Println(largestComponentSize([]int{83, 99, 39, 11, 19, 30, 31}))
	fmt.Println(largestComponentSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
