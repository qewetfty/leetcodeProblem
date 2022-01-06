package array

//Anti-theft security devices are activated inside a bank. You are given a 0-
//indexed binary string array bank representing the floor plan of the bank, which is
//an m x n 2D matrix. bank[i] represents the iᵗʰ row, consisting of '0's and '1's.
// '0' means the cell is empty, while'1' means the cell has a security device.
//
// There is one laser beam between any two security devices if both conditions
//are met:
//
//
// The two devices are located on two different rows: r1 and r2, where r1 < r2.
//
// For each row i where r1 < i < r2, there are no security devices in the iᵗʰ
//row.
//
//
// Laser beams are independent, i.e., one beam does not interfere nor join with
//another.
//
// Return the total number of laser beams in the bank.
//
//
// Example 1:
//
//
//Input: bank = ["011001","000000","010100","001000"]
//Output: 8
//Explanation: Between each of the following device pairs, there is one beam.
//In total, there are 8 beams:
// * bank[0][1] -- bank[2][1]
// * bank[0][1] -- bank[2][3]
// * bank[0][2] -- bank[2][1]
// * bank[0][2] -- bank[2][3]
// * bank[0][5] -- bank[2][1]
// * bank[0][5] -- bank[2][3]
// * bank[2][1] -- bank[3][2]
// * bank[2][3] -- bank[3][2]
//Note that there is no beam between any device on the 0ᵗʰ row with any on the 3
//ʳᵈ row.
//This is because the 2ⁿᵈ row contains security devices, which breaks the
//second condition.
//
//
// Example 2:
//
//
//Input: bank = ["000","111","000"]
//Output: 0
//Explanation: There does not exist two devices located on two different rows.
//
//
//
// Constraints:
//
//
// m == bank.length
// n == bank[i].length
// 1 <= m, n <= 500
// bank[i][j] is either '0' or '1'.
//
// Related Topics Array Math String Matrix 👍 125 👎 8

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfBeams(bank []string) int {
	numList := make([]int, 0)
	for _, s := range bank {
		num := 0
		for _, char := range s {
			if char == '1' {
				num++
			}
		}
		if num > 0 {
			numList = append(numList, num)
		}
	}
	l := len(numList)
	if l < 2 {
		return 0
	}
	result := 0
	for i := 1; i < l; i++ {
		result = result + numList[i-1]*numList[i]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
