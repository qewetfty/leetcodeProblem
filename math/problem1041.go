package math

import "fmt"

// On an infinite plane, a robot initially stands at (0, 0) and faces north.
// The robot can receive one of three instructions:
//	"G": go straight 1 unit;
//	"L": turn 90 degrees to the left;
//	"R": turn 90 degress to the right.
// The robot performs the instructions given in order, and repeats them forever.
// Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.
//	Example 1:
//		Input: "GGLLGG"
//		Output: true
//		Explanation:
//			The robot moves from (0,0) to (0,2), turns 180 degrees, and then returns to (0,0).
//			When repeating these instructions, the robot remains in the circle of radius 2 centered at the origin.
//	Example 2:
//		Input: "GG"
//		Output: false
//		Explanation:
//			The robot moves north indefinitely.
//	Example 3:
//		Input: "GL"
//		Output: true
//		Explanation:
//			The robot moves from (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ...
//	Note:
//		1 <= instructions.length <= 100
//		instructions[i] is in {'G', 'L', 'R'}

func isRobotBounded(instructions string) bool {
	move, flag := make([]int, 4), 0
	for i := 0; i < 4; i++ {
		for _, char := range instructions {
			if char == 'G' {
				move[flag]++
			} else if char == 'L' {
				flag += 3
			} else if char == 'R' {
				flag += 1
			}
			flag = flag % 4
		}
	}
	if move[0] == move[2] && move[1] == move[3] {
		return true
	}
	return false
}

func testProblem1041() {
	fmt.Println(isRobotBounded("GGLLGG"))
	fmt.Println(isRobotBounded("GG"))
	fmt.Println(isRobotBounded("GL"))
}
