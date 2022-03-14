package utils

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Gcd 求解两个数的最大公约数
func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return Gcd(y%x, x)
}
