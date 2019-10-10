package math

import "fmt"

// Write an algorithm to determine if a number is "happy".

// A happy number is a number defined by the following process:
// Starting with any positive integer, replace the number by the sum of the squares of its digits,
// and repeat the process until the number equals 1 (where it will stay),
// or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy numbers.

func test() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	numberMap := make(map[int]int, 0)
	for {
		sum := 0
		for n != 0 {
			num := n % 10
			n = n / 10
			sum += num * num
		}
		if sum == 1 {
			return true
		}
		if _, ok := numberMap[sum]; ok {
			return false
		} else {
			numberMap[sum] = 0
		}
		n = sum
	}
}
