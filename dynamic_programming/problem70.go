package dynamic_programming

import "fmt"

// You are climbing a stair case. It takes n steps to reach to the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Note: Given n will be a positive integer.

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-2) + climbStairs(n-1)
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first := 1
	second := 2
	for i := 2; i < n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func testProblem70() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(1000))
}
