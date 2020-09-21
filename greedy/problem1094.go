package greedy

import (
	"fmt"
	"sort"
)

// You are driving a vehicle that has capacity empty seats initially available for passengers.  The vehicle only drives
// east (ie. it cannot turn around and drive west.)
// Given a list of trips, trip[i] = [num_passengers, start_location, end_location] contains information about the i-th
// trip: the number of passengers that must be picked up, and the locations to pick them up and drop them off.
// The locations are given as the number of kilometers due east from your vehicle's initial location.
// Return true if and only if it is possible to pick up and drop off all passengers for all the given trips.
//	Example 1:
//		Input: trips = [[2,1,5],[3,3,7]], capacity = 4
//		Output: false
//	Example 2:
//		Input: trips = [[2,1,5],[3,3,7]], capacity = 5
//		Output: true
//	Example 3:
//		Input: trips = [[2,1,5],[3,5,7]], capacity = 3
//		Output: true
//	Example 4:
//		Input: trips = [[3,2,7],[3,7,9],[8,3,9]], capacity = 11
//		Output: true
//	Constraints:
//		trips.length <= 1000
//		trips[i].length == 3
//		1 <= trips[i][0] <= 100
//		0 <= trips[i][1] < trips[i][2] <= 1000
//		1 <= capacity <= 100000

// 记录负载
func carPooling2(trips [][]int, capacity int) bool {
	l := len(trips)
	if l == 0 {
		return true
	}
	locations := make([]int, 1001)
	sort.SliceStable(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	for _, trip := range trips {
		passenger := trip[0]
		for i := trip[1]; i < trip[2]; i++ {
			locations[i] = locations[i] + passenger
			if locations[i] > capacity {
				return false
			}
		}
	}
	return true
}

// 只记录容量变化，查询容量是否有小于0的情况
func carPooling(trips [][]int, capacity int) bool {
	capacityChange := make([]int, 1001)
	for _, trip := range trips {
		capacityChange[trip[1]] -= trip[0]
		capacityChange[trip[2]] += trip[0]
	}
	for i := 0; i < 1001; i++ {
		capacity += capacityChange[i]
		if capacity < 0 {
			return false
		}
	}
	return true
}

func testProblem1094() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
	fmt.Println(carPooling([][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11))
}
