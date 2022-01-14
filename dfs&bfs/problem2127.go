package dfs_bfs

//A company is organizing a meeting and has a list of n employees, waiting to
//be invited. They have arranged for a large circular table, capable of seating any
//number of employees.
//
// The employees are numbered from 0 to n - 1. Each employee has a favorite
//person and they will attend the meeting only if they can sit next to their favorite
//person at the table. The favorite person of an employee is not themself.
//
// Given a 0-indexed integer array favorite, where favorite[i] denotes the
//favorite person of the iᵗʰ employee, return the maximum number of employees that can
//be invited to the meeting.
//
//
// Example 1:
//
//
//Input: favorite = [2,2,1,2]
//Output: 3
//Explanation:
//The above figure shows how the company can invite employees 0, 1, and 2, and
//seat them at the round table.
//All employees cannot be invited because employee 2 cannot sit beside
//employees 0, 1, and 3, simultaneously.
//Note that the company can also invite employees 1, 2, and 3, and give them
//their desired seats.
//The maximum number of employees that can be invited to the meeting is 3.
//
//
// Example 2:
//
//
//Input: favorite = [1,2,0]
//Output: 3
//Explanation:
//Each employee is the favorite person of at least one other employee, and the
//only way the company can invite them is if they invite every employee.
//The seating arrangement will be the same as that in the figure given in
//example 1:
//- Employee 0 will sit between employees 2 and 1.
//- Employee 1 will sit between employees 0 and 2.
//- Employee 2 will sit between employees 1 and 0.
//The maximum number of employees that can be invited to the meeting is 3.
//
//
// Example 3:
//
//
//Input: favorite = [3,0,1,4,1]
//Output: 4
//Explanation:
//The above figure shows how the company will invite employees 0, 1, 3, and 4,
//and seat them at the round table.
//Employee 2 cannot be invited because the two spots next to their favorite
//employee 1 are taken.
//So the company leaves them out of the meeting.
//The maximum number of employees that can be invited to the meeting is 4.
//
//
//
// Constraints:
//
//
// n == favorite.length
// 2 <= n <= 10⁵
// 0 <= favorite[i] <= n - 1
// favorite[i] != i
//
// Related Topics Depth-First Search Graph Topological Sort 👍 310 👎 2

//leetcode submit region begin(Prohibit modification and deletion)
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	indeg := make([]int, n)
	for _, f := range favorite {
		indeg[f]++
	}
	used, dp := make([]bool, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		used[u] = true
		queue = queue[1:]
		v := favorite[u]
		dp[v] = Max(dp[v], dp[u]+1)
		indeg[v]--
		if indeg[v] == 0 {
			queue = append(queue, v)
		}
	}
	ring, total := 0, 0
	for i := 0; i < n; i++ {
		if !used[i] {
			j := favorite[i]
			if favorite[j] == i {
				total = total + dp[i] + dp[j]
				used[i], used[j] = true, true
			} else {
				u, cnt := i, 0
				for {
					cnt++
					u = favorite[u]
					used[u] = true
					if u == i {
						break
					}
				}
				ring = Max(ring, cnt)
			}
		}
	}
	return Max(ring, total)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
