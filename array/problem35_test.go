package array

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	input1 := []int{1, 3, 5, 6}
	target1 := 5
	fmt.Println(SearchInsert(input1, target1))
	target2 := 2
	fmt.Println(SearchInsert(input1, target2))
	target3 := 7
	fmt.Println(SearchInsert(input1, target3))
	target4 := 0
	fmt.Println(SearchInsert(input1, target4))
	input2 := []int{1, 3}
	fmt.Println(SearchInsert(input2, target4))
	fmt.Println(SearchInsert(input2, target2))
}
