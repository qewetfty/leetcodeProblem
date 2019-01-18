package dynamic_programming

import (
	"fmt"
	"math"
)

//Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.
//You have the following 3 operations permitted on a word:
//Insert a character
//Delete a character
//Replace a character

// string, dynamic programming

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	if l1*l2 == 0 {
		return l1 + l2
	}
	//init 2-dim slice
	distance := make([][]int, l1+1)
	for k := 0; k < l1+1; k++ {
		distance[k] = make([]int, l2+1)
	}
	for i := 0; i < l1+1; i++ {
		distance[i][0] = i
	}
	for j := 0; j < l2+1; j++ {
		distance[0][j] = j
	}
	s1 := []byte(word1)
	s2 := []byte(word2)
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			left := distance[i-1][j] + 1
			down := distance[i][j-1] + 1
			leftDown := distance[i-1][j-1]
			if s1[i-1] != s2[j-1] {
				leftDown += 1
			}
			distance[i][j] = int(math.Min(float64(left), math.Min(float64(down), float64(leftDown))))
		}
	}
	return distance[l1][l2]
}

func testProblem72() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
	word3 := "intention"
	word4 := "execution"
	fmt.Println(minDistance(word3, word4))
	word5 := "pneumonoultramicroscopicsilicovolcanoconiosis"
	word6 := "ultramicroscopically"
	fmt.Println(minDistance(word5, word6))
}
