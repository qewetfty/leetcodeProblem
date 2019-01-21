package math

import "fmt"

// Given a positive integer N, find and return the longest distance between
// two consecutive 1's in the binary representation of N.

//If there aren't two consecutive 1's, return 0.

func binaryGap(N int) int {
	max := 0
	interval := 0
	for ; N%2 == 0; N = N >> 1 {
		continue
	}
	if N == 1 {
		return max
	}
	interval = 1
	max = 1
	for ; N > 0; N = N >> 1 {
		if N%2 == 0 {
			interval++
			continue
		}
		max = maxInt(max, interval)
		interval = 1
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testProblem868() {
	fmt.Println(binaryGap(7))
	fmt.Println(binaryGap(6))
	fmt.Println(binaryGap(5))
	fmt.Println(binaryGap(8))
}
