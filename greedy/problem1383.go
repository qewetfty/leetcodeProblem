package greedy

import (
	"container/heap"
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// You are given two integers n and k and two integer arrays speed and efficiency
// both of length n. There are n engineers numbered from 1 to n. speed[i] and
// efficiency[i] represent the speed and efficiency of the ith engineer
// respectively.
// Choose at most k different engineers out of the n engineers to form a team with the maximum performance.
// The performance of a team is the sum of their engineers' speeds multiplied by the minimum efficiency among their engineers.
// Return the maximum performance of this team. Since the answer can be a huge number, return it modulo 109 + 7.
//	Example 1:
//		Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
//		Output: 60
//		Explanation:
//			We have the maximum performance of the team by selecting engineer 2 (with speed=10 and efficiency=4) and engineer 5 (with speed=5 and efficiency=7). That is, performance = (10 + 5) * min(4, 7) = 60.
//	Example 2:
//		Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
//		Output: 68
//		Explanation:
//			This is the same example as the first but k = 3. We can select engineer 1, engineer 2 and engineer 5 to get the maximum performance of the team. That is, performance = (2 + 10 + 5) * min(5, 4, 7) = 68.
//	Example 3:
//		Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
//		Output: 72
//	Constraints:
//		1 <= <= k <= n <= 105
//		speed.length == n
//		efficiency.length == n
//		1 <= speed[i] <= 105
//		1 <= efficiency[i] <= 108

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	// 效率倒序
	efficiencyList := make([]worker, 0)
	for i := 0; i < n; i++ {
		efficiencyList = append(efficiencyList, worker{speed: speed[i], efficiency: efficiency[i]})
	}
	sort.Slice(efficiencyList, func(i, j int) bool {
		return efficiencyList[i].efficiency < efficiencyList[j].efficiency
	})
	// 速度的最小堆
	speedList := new(workers)
	sumSpeed := 0
	result := 0
	for i := n - 1; i >= n-k; i-- {
		heap.Push(speedList, efficiencyList[i])
		sumSpeed += efficiencyList[i].speed
		result = utils.Max(sumSpeed*efficiencyList[i].efficiency, result)
	}
	for i := n - k - 1; i >= 0; i-- {
		sumSpeed += efficiencyList[i].speed
		heap.Push(speedList, efficiencyList[i])
		top := heap.Pop(speedList).(worker)
		sumSpeed -= top.speed
		result = utils.Max(result, sumSpeed*efficiencyList[i].efficiency)
	}
	return result % 1000000007
}

type worker struct {
	speed      int
	efficiency int
}

type workers []worker

func (w workers) Len() int {
	return len(w)
}

func (w workers) Less(i, j int) bool {
	return w[i].speed < w[j].speed
}

func (w workers) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w *workers) Push(x interface{}) {
	*w = append(*w, x.(worker))
}

func (w *workers) Pop() interface{} {
	old := *w
	n := len(old)
	x := old[n-1]
	*w = old[:n-1]
	return x
}

func testProblem1383() {
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2))
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3))
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4))
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 5))
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 6))
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 1))
}
