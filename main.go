package main

import (
	"fmt"
	"math/big"
)

func countTexts(pressedKeys string) int {
	l := len(pressedKeys)
	char, count := pressedKeys[0], 1
	result := big.NewInt(1)
	for i := 1; i < l; i++ {
		curChar := pressedKeys[i]
		if curChar == char {
			count++
		} else {
			curCount := getCount(count)
			result = result.Mul(result, big.NewInt(int64(curCount)))
			char, count = curChar, 1
		}
	}
	curCount := getCount(count)
	result = result.Mul(result, big.NewInt(int64(curCount)))
	result = result.Mod(result, big.NewInt(1000000007))
	return int(result.Int64())
}

func getCount(n int) int64 {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	var (
		x, y, z int64
	)
	x, y, z = 1, 2, 4
	for i := 4; i <= n; i++ {
		x, y, z = y, z, x+y+z
	}
	return z
}

func main() {
	fmt.Println(getCount(10000))
	fmt.Println(countTexts("2"))
	fmt.Println(countTexts("22"))
	fmt.Println(countTexts("222"))
	fmt.Println(countTexts("2222"))
	fmt.Println(countTexts("22222"))
	fmt.Println(countTexts("222222"))
}
