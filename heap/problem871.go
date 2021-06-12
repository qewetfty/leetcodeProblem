package heap

import (
	"container/heap"
	"fmt"
	"github.com/leetcodeProblem/data"
	"github.com/leetcodeProblem/utils"
)

// A car travels from a starting position to a destination which is target miles
// east of the starting position.
// Along the way, there are gas stations. Each station[i] represents a gas
// station that is station[i][0] miles east of the starting position, and has
// station[i][1] liters of gas.
// The car starts with an infinite tank of gas, which initially has startFuel
// liters of fuel in it. It uses 1 liter of gas per 1 mile that it drives.
// When the car reaches a gas station, it may stop and refuel, transferring all the gas from the station into the car.
// What is the least number of refueling stops the car must make in order to
// reach its destination? If it cannot reach the destination, return -1.
// Note that if the car reaches a gas station with 0 fuel left, the car can still
// refuel there. If the car reaches the destination with 0 fuel left, it is still
// considered to have arrived.
//	Example 1:
//		Input: target = 1, startFuel = 1, stations = []
//		Output: 0
//		Explanation: We can reach the target without refueling.
//	Example 2:
//		Input: target = 100, startFuel = 1, stations = [[10,100]]
//		Output: -1
//		Explanation: We can't reach the target (or even the first gas station).
//	Example 3:
//		Input: target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
//		Output: 2
//		Explanation:
//			We start with 10 liters of fuel.
//			We drive to position 10, expending 10 liters of fuel.  We refuel from 0 liters to 60 liters of gas.
//			Then, we drive from position 10 to position 60 (expending 50 liters of fuel),
//			and refuel from 10 liters to 50 liters of gas.  We then drive to and reach the target.
//			We made 2 refueling stops along the way, so we return 2.
//	Note:
//		1 <= target, startFuel, stations[i][1] <= 10^9
//		0 <= stations.length <= 500
//		0 < stations[0][0] < stations[1][0] < ... < stations[stations.length-1][0] < target

// dp解法，定义dp[i]为加i次油可以到达的最远位置。
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	l := len(stations)
	dp := make([]int, l+1)
	dp[0] = startFuel
	for i := 0; i < l; i++ {
		for t := i; t >= 0; t-- {
			if dp[t] >= stations[i][0] {
				dp[t+1] = utils.Max(dp[t+1], dp[t]+stations[i][1])
			}
		}
	}
	for i, miles := range dp {
		if miles >= target {
			return i
		}
	}
	return -1
}

// heap解法，每次贪心获取最大油量的station，计算跑的最远的地方
func minRefuelStops2(target int, tank int, stations [][]int) int {
	maxGasStation := new(data.IntHeapDesc)
	result, prev := 0, 0
	for _, station := range stations {
		location, capcity := station[0], station[1]
		tank -= location - prev
		for maxGasStation.Len() > 0 && tank < 0 {
			tank = tank + heap.Pop(maxGasStation).(int)
			result++
		}
		if tank < 0 {
			return -1
		}
		heap.Push(maxGasStation, capcity)
		prev = location
	}
	tank -= target - prev
	for maxGasStation.Len() > 0 && tank < 0 {
		tank = tank + heap.Pop(maxGasStation).(int)
		result++
	}
	if tank < 0 {
		return -1
	}
	return result
}

func testProblem871() {
	fmt.Println(minRefuelStops2(1, 1, [][]int{}))
	fmt.Println(minRefuelStops2(100, 1, [][]int{{10, 100}}))
	fmt.Println(minRefuelStops2(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}
