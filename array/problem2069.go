package array

//A width x height grid is on an XY-plane with the bottom-left cell at (0, 0)
//and the top-right cell at (width - 1, height - 1). The grid is aligned with the
//four cardinal directions ("North", "East", "South", and "West"). A robot is
//initially at cell (0, 0) facing direction "East".
//
// The robot can be instructed to move for a specific number of steps. For each
//step, it does the following.
//
//
// Attempts to move forward one cell in the direction it is facing.
// If the cell the robot is moving to is out of bounds, the robot instead turns
//90 degrees counterclockwise and retries the step.
//
//
// After the robot finishes moving the number of steps required, it stops and
//awaits the next instruction.
//
// Implement the Robot class:
//
//
// Robot(int width, int height) Initializes the width x height grid with the
//robot at (0, 0) facing "East".
// void step(int num) Instructs the robot to move forward num steps.
// int[] getPos() Returns the current cell the robot is at, as an array of
//length 2, [x, y].
// String getDir() Returns the current direction of the robot, "North", "East",
//"South", or "West".
//
//
//
// Example 1:
//
//
//Input
//["Robot", "move", "move", "getPos", "getDir", "move", "move", "move",
//"getPos", "getDir"]
//[[6, 3], [2], [2], [], [], [2], [1], [4], [], []]
//Output
//[null, null, null, [4, 0], "East", null, null, null, [1, 2], "West"]
//
//Explanation
//Robot robot = new Robot(6, 3); // Initialize the grid and the robot at (0, 0)
//facing East.
//robot.move(2);  // It moves two steps East to (2, 0), and faces East.
//robot.move(2);  // It moves two steps East to (4, 0), and faces East.
//robot.getPos(); // return [4, 0]
//robot.getDir(); // return "East"
//robot.move(2);  // It moves one step East to (5, 0), and faces East.
//                // Moving the next step East would be out of bounds, so it
//turns and faces North.
//                // Then, it moves one step North to (5, 1), and faces North.
//robot.move(1);  // It moves one step North to (5, 2), and faces North (not
//West).
//robot.move(4);  // Moving the next step North would be out of bounds, so it
//turns and faces West.
//                // Then, it moves four steps West to (1, 2), and faces West.
//robot.getPos(); // return [1, 2]
//robot.getDir(); // return "West"
//
//
//
//
// Constraints:
//
//
// 2 <= width, height <= 100
// 1 <= num <= 10⁵
// At most 10⁴ calls in total will be made to step, getPos, and getDir.
//
// Related Topics Design Simulation 👍 75 👎 186

//leetcode submit region begin(Prohibit modification and deletion)
type Robot struct {
	index        int
	moved        bool
	position     [][]int
	direction    []int
	directionMap map[int]string
}

func Constructor2069(width int, height int) Robot {
	pos, direction := make([][]int, 0), make([]int, 0)
	for i := 0; i < width; i++ {
		pos = append(pos, []int{i, 0})
		direction = append(direction, 0)
	}
	for i := 1; i < height; i++ {
		pos = append(pos, []int{width - 1, i})
		direction = append(direction, 1)
	}
	for i := width - 2; i >= 0; i-- {
		pos = append(pos, []int{i, height - 1})
		direction = append(direction, 2)
	}
	for i := height - 2; i > 0; i-- {
		pos = append(pos, []int{0, i})
		direction = append(direction, 3)
	}
	direction[0] = 3
	return Robot{
		index:        0,
		moved:        false,
		position:     pos,
		direction:    direction,
		directionMap: map[int]string{0: "East", 1: "North", 2: "West", 3: "South"},
	}
}

func (this *Robot) Step(num int) {
	this.moved = true
	this.index = (this.index + num) % len(this.position)
}

func (this *Robot) GetPos() []int {
	return this.position[this.index]
}

func (this *Robot) GetDir() string {
	if !this.moved {
		return "East"
	}
	return this.directionMap[this.direction[this.index]]
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
//leetcode submit region end(Prohibit modification and deletion)
