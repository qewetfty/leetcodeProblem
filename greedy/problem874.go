package greedy

import (
	"fmt"
	"math"
)

// A robot on an infinite grid starts at point (0, 0) and faces north.  The robot can receive one of three possible types of commands:
//	-2: turn left 90 degrees
//	-1: turn right 90 degrees
//	1 <= x <= 9: move forward x units
// Some of the grid squares are obstacles.
// The i-th obstacle is at grid point (obstacles[i][0], obstacles[i][1])
// If the robot would try to move onto them, the robot stays on the previous grid square instead
// (but still continues following the rest of the route.)
// Return the square of the maximum Euclidean distance that the robot will be from the origin.
//	Example 1:
//		Input: commands = [4,-1,3], obstacles = []
//		Output: 25
//		Explanation: robot will go to (3, 4)
//	Example 2:
//		Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
//		Output: 65
//		Explanation: robot will be stuck at (1, 4) before turning left and going to (1, 8)
//	Note:
//		0 <= commands.length <= 10000
//		0 <= obstacles.length <= 10000
//		-30000 <= obstacle[i][0] <= 30000
//		-30000 <= obstacle[i][1] <= 30000
//	The answer is guaranteed to be less than 2 ^ 31.

type point struct {
	x int
	y int
}

var direction = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func robotSim(commands []int, obstacles [][]int) int {
	res := 0
	if len(commands) == 0 {
		return res
	}
	obMap := make(map[point]bool)
	for _, obstacle := range obstacles {
		obMap[point{x: obstacle[0], y: obstacle[1]}] = true
	}
	var (
		dir = 0
		x   = 0
		y   = 0
	)
	for _, command := range commands {
		if command == -2 {
			dir = (dir + 4 - 1) % 4
			continue
		} else if command == -1 {
			dir = (dir + 1) % 4
			continue
		}
		for i := 0; i < command; i++ {
			newX := x + direction[dir][0]
			newY := y + direction[dir][1]
			if _, ok := obMap[point{x: newX, y: newY}]; ok {
				break
			}
			x, y = newX, newY
		}
		res = int(math.Max(float64(res), float64(x*x+y*y)))
	}
	return res
}

func testProblem874() {
	fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}))
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}
