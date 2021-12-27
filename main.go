package main

import (
	"fmt"
	"sort"
)

//You are given an integer n indicating there are n people numbered from 0 to n
//- 1. You are also given a 0-indexed 2D integer array meetings where meetings[i]
//= [xi, yi, timei] indicates that person xi and person yi have a meeting at
//timei. A person may attend multiple meetings at the same time. Finally, you are
//given an integer firstPerson.
//
// Person 0 has a secret and initially shares the secret with a person
//firstPerson at time 0. This secret is then shared every time a meeting takes place with
//a person that has the secret. More formally, for every meeting, if a person xi
//has the secret at timei, then they will share the secret with person yi, and vice
//versa.
//
// The secrets are shared instantaneously. That is, a person may receive the
//secret and share it with people in other meetings within the same time frame.
//
// Return a list of all the people that have the secret after all the meetings
//have taken place. You may return the answer in any order.
//
//
// Example 1:
//
//
//Input: n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1
//Output: [0,1,2,3,5]
//Explanation:
//At time 0, person 0 shares the secret with person 1.
//At time 5, person 1 shares the secret with person 2.
//At time 8, person 2 shares the secret with person 3.
//At time 10, person 1 shares the secret with person 5.​​​​
//Thus, people 0, 1, 2, 3, and 5 know the secret after all the meetings.
//
//
// Example 2:
//
//
//Input: n = 4, meetings = [[3,1,3],[1,2,2],[0,3,3]], firstPerson = 3
//Output: [0,1,3]
//Explanation:
//At time 0, person 0 shares the secret with person 3.
//At time 2, neither person 1 nor person 2 know the secret.
//At time 3, person 3 shares the secret with person 0 and person 1.
//Thus, people 0, 1, and 3 know the secret after all the meetings.
//
//
// Example 3:
//
//
//Input: n = 5, meetings = [[3,4,2],[1,2,1],[2,3,1]], firstPerson = 1
//Output: [0,1,2,3,4]
//Explanation:
//At time 0, person 0 shares the secret with person 1.
//At time 1, person 1 shares the secret with person 2, and person 2 shares the
//secret with person 3.
//Note that person 2 can share the secret at the same time as receiving it.
//At time 2, person 3 shares the secret with person 4.
//Thus, people 0, 1, 2, 3, and 4 know the secret after all the meetings.
//
//
// Example 4:
//
//
//Input: n = 6, meetings = [[0,2,1],[1,3,1],[4,5,1]], firstPerson = 1
//Output: [0,1,2,3]
//Explanation:
//At time 0, person 0 shares the secret with person 1.
//At time 1, person 0 shares the secret with person 2, and person 1 shares the
//secret with person 3.
//Thus, people 0, 1, 2, and 3 know the secret after all the meetings.
//
//
//
// Constraints:
//
//
// 2 <= n <= 10⁵
// 1 <= meetings.length <= 10⁵
// meetings[i].length == 3
// 0 <= xi, yi <= n - 1
// xi != yi
// 1 <= timei <= 10⁵
// 1 <= firstPerson <= n - 1
//
// Related Topics Depth-First Search Breadth-First Search Union Find Graph
//Sorting 👍 225 👎 14

//leetcode submit region begin(Prohibit modification and deletion)
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	u := NewUnionFind(n)
	// 对meeting按照时间排序，按照时间从小到大开始执行Union操作
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	l := len(meetings)
	u.Union(0, firstPerson)
	for i := 0; i < l; i++ {
		j := i + 1
		for j < l && meetings[i][2] == meetings[j][2] {
			j++
		}
		// 先把同一时间开会的连起来
		for k := i; k < j; k++ {
			u.Union(meetings[k][0], meetings[k][1])
		}
		// 把没有和0连通的进行isolate操作
		for k := i; k < j; k++ {
			if u.Parent(meetings[k][0]) != u.Parent(0) {
				u.Isolate(meetings[k][0])
				u.Isolate(meetings[k][1])
			}
		}
		i = j - 1
	}
	// 计算结果，即获取所有和0号都是同一parent的编号
	result, zeroParent := []int{0}, u.Parent(0)
	for i := 1; i < n; i++ {
		if u.Parent(i) == zeroParent {
			result = append(result, i)
		}
	}
	return result
}

type unionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) unionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return unionFind{
		parent: parent,
		count:  n,
	}
}

func (u unionFind) Parent(n int) int {
	root, i := n, n
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u unionFind) Union(i, j int) {
	m, n := u.Parent(i), u.Parent(j)
	if m == n {
		return
	}
	u.parent[m] = n
	u.count--
}

func (u unionFind) Count() int {
	return u.count
}

func (u unionFind) Isolate(i int) {
	if u.Parent(i) == i {
		return
	}
	u.parent[i] = i
	u.count++
}

func main() {
	fmt.Println(1 ^ 1)
	fmt.Println(1 ^ 0)
	fmt.Println(0 ^ 1)
	fmt.Println(0 ^ 0)

}
